package softcanvas

import (
	"github.com/iamGreedy/canvas"
	"github.com/iamGreedy/canvas/claa"
	"image"
	"image/color"
	"image/draw"
	"math"
	"sync"
	"sync/atomic"
)

const (
	pixR = 0
	pixG = 1
	pixB = 2
	pixA = 3
	pixL = 4
)

type Software struct {
	//
	size      image.Point
	pix       []uint8
	mix       canvas.MixOperation
	precision claa.Precision
	// Mutex for each lines
	linelck []sync.Mutex
	// Current working group lock
	works sync.WaitGroup
	// Buffer for Query
	queryon     int32
	querybuffer []q
	queryblck   sync.Mutex
	// *claa.CLAA Pool
	pool *sync.Pool
}
type q struct {
	s canvas.Shader
	p *canvas.Path
	t *canvas.Transform
}

func NewSoftware(w, h int, options ...canvas.Option) *Software {
	res := new(Software)
	res.size = image.Pt(w, h)
	res.precision = claa.X16
	res.setup()
	for _, o := range options {
		if o.OptionType()&canvas.Init == canvas.Init {
			switch t := o.(type) {
			case canvas.Clear:
				res.Option(t)
			case canvas.MixOperation:
				res.Option(t)
			case OptionCLAAPrecision:
				res.Option(t)

			}
		}
	}
	return res
}

// Surface interface
func (s *Software) Support(opt ...canvas.Option) bool {
	for _, o := range opt {
		if !s.support(o) {
			return false
		}
	}
	return true
}
// Surface interface
func (s *Software) support(opt canvas.Option) bool {
	switch opt.(type) {
	case canvas.Width:
		return true
	case *canvas.Width:
		return true
	case canvas.Height:
		return true
	case *canvas.Height:
		return true
	case canvas.Size:
		return true
	case *canvas.Size:
		return true
	case *canvas.Driver:
		return true
	case canvas.Clear:
		return true
	case canvas.FrameBuffer:
		return true
	case *canvas.FrameBuffer:
		return true
	//
	case canvas.UseBuffer:
		return true
	case *canvas.UseBuffer:
		return true
	case *canvas.BufferLength:
		return true
	case canvas.CloseBuffer:
		return true
	case *canvas.CloseBuffer:
		return true
	case canvas.FlushBuffer:
		return true
	case *canvas.FlushBuffer:
		return true
	case canvas.WaitFlush:
		return true
	case *canvas.WaitFlush:
		return true
	}
	return false
}
// Surface interface
func (s *Software) Option(opt canvas.Option) canvas.Option {
	switch o := opt.(type) {
	case canvas.Width:
		s.size.X = int(o)
		s.setup()
		return canvas.Width(s.size.X)
	case *canvas.Width:
		if o == nil {
			o = new(canvas.Width)
		}
		*o = canvas.Width(s.size.X)
		return o
	case canvas.Height:
		s.size.Y = int(o)
		s.setup()
		return canvas.Height(s.size.Y)
	case *canvas.Height:
		if o == nil {
			o = new(canvas.Height)
		}
		*o = canvas.Height(s.size.Y)
		return o
	case canvas.Size:
		s.size = image.Point(o)
		s.setup()
		return canvas.Size(s.size)
	case *canvas.Size:
		if o == nil {
			o = new(canvas.Size)
		}
		*o = canvas.Size(s.size)
		return o
	case *canvas.Driver:
		if o == nil {
			o = new(canvas.Driver)
		}
		*o = Driver
		return o
	case canvas.Clear:
		s.Clear(color.RGBA(o))
	case canvas.FrameBuffer:
		temp := &image.RGBA{
			Pix:    s.pix,
			Stride: s.size.X * pixL,
			Rect:   image.Rect(0, 0, s.size.X, s.size.Y),
		}
		orfr := image.RGBA(o)
		draw.Draw(temp, temp.Rect, &orfr, image.ZP, draw.Src)
		return o
	case *canvas.FrameBuffer:
		if o == nil {
			o = new(canvas.FrameBuffer)
		}
		o.Stride = pixL * s.size.X
		o.Rect = image.Rect(0, 0, s.size.X, s.size.Y)
		o.Pix = make([]uint8, len(s.pix))
		copy(o.Pix, s.pix)
		return o
	case canvas.UseBuffer:
		atomic.StoreInt32(&s.queryon, 1)
	case *canvas.UseBuffer:
		atomic.StoreInt32(&s.queryon, 1)
	case *canvas.BufferLength:
		*o = canvas.BufferLength(len(s.querybuffer))
	case canvas.CloseBuffer:
		atomic.StoreInt32(&s.queryon, 0)
		s.querybuffer = nil
	case *canvas.CloseBuffer:
		atomic.StoreInt32(&s.queryon, 0)
		s.querybuffer = nil
	case canvas.FlushBuffer:
		if atomic.CompareAndSwapInt32(&s.queryon, 1, 1) {
			s.works.Add(len(s.querybuffer))
			s.queryblck.Lock()
			for _, value := range s.querybuffer {
				go func(value q) {
					s.query(value.p, value.s, value.t)
					s.works.Done()
				}(value)
			}
			s.queryblck.Unlock()
			s.querybuffer = s.querybuffer[:0]
		}
		//
	case *canvas.FlushBuffer:
		if atomic.CompareAndSwapInt32(&s.queryon, 1, 1) {
			s.works.Add(len(s.querybuffer))
			s.queryblck.Lock()
			for _, value := range s.querybuffer {
				go func(value q) {
					s.query(value.p, value.s, value.t)
					s.works.Done()
				}(value)
			}
			s.queryblck.Unlock()
			s.querybuffer = s.querybuffer[:0]
		}
		//
	case canvas.WaitFlush:
		s.works.Wait()
	case *canvas.WaitFlush:
		s.works.Wait()
	case OptionCLAAPrecision:
		s.precision = claa.Precision(o)
		return OptionCLAAPrecision(s.precision)
	case *OptionCLAAPrecision:
		if o == nil {
			o = new(OptionCLAAPrecision)
		}
		*o = OptionCLAAPrecision(s.precision)
		return o
	}
	return nil
}
// Surface interface
func (s *Software) Query(query *canvas.Path, shader canvas.Shader, transform *canvas.Transform) error {
	if atomic.CompareAndSwapInt32(&s.queryon, 1, 1) {
		s.queryblck.Lock()
		s.querybuffer = append(s.querybuffer, q{
			s: shader,
			p: query,
			t: transform,
		})
		s.queryblck.Unlock()
	} else {
		s.query(query, shader, transform)
	}
	return nil
}
// Surface interface
func (s *Software) query(query *canvas.Path, shader canvas.Shader, transform *canvas.Transform) {
	ws := s.pool.Get().(*claa.CLAA)
	defer s.pool.Put(ws)
	stencil := make([]uint8, s.size.X*s.size.Y)
	//
	ws.Data(claa.RAW, query.Data...)
	ws.Result(stencil, s.size.X, 1, 0)
	switch shd := shader.(type) {
	case *canvas.ColorShader:
		for y := 0; y < s.size.Y; y++ {
			for x := 0; x < s.size.X; x++ {
				offset := s.size.X*4*y + x*4
				if s.mix == canvas.Over {
					mixRGBA32Over(s.pix[offset:offset+pixL], shd.R, shd.G, shd.B, shd.A, stencil[offset/4])
				} else {
					mixRGBA32Src(s.pix[offset:offset+pixL], shd.R, shd.G, shd.B, shd.A, stencil[offset/4])
				}
			}
		}
	}
	ws.Clear()
}

// Convenient Method
func (s *Software) Clear(c color.RGBA) {
	s.works.Add(1)
	defer s.works.Done()
	//
	wg := new(sync.WaitGroup)
	wg.Add(s.size.Y)
	for y := 0; y < s.size.Y; y++ {
		go func(y int) {
			for x := 0; x < s.size.X; x++ {
				offset := s.size.X*4*y + x*pixL
				s.pix[offset+pixR] = c.R
				s.pix[offset+pixG] = c.G
				s.pix[offset+pixB] = c.B
				s.pix[offset+pixA] = c.A
			}
			wg.Done()
		}(y)
	}
	wg.Wait()
}




// private usage
func (s *Software) setup() {
	s.pix = make([]uint8, s.size.X*s.size.Y*pixL)
	s.linelck = make([]sync.Mutex, s.size.Y)
	s.pool = &sync.Pool{
		New: func() interface{} {
			return claa.NewCLAA(s.precision, s.size.X, s.size.Y)
		},
	}
}
// private usage
func intense(r, g, b, a uint8, i uint8) (or, og, ob, oa uint8) {
	return uint8(uint16(r) * uint16(i) / math.MaxUint8), uint8(uint16(g) * uint16(i) / math.MaxUint8), uint8(uint16(b) * uint16(i) / math.MaxUint8), uint8(uint16(a) * uint16(i) / math.MaxUint8)
}
// private usage
func mixRGBA32Over(dst []uint8, src ...uint8) {
	src[pixR], src[pixG], src[pixB], src[pixA] = intense(src[pixR], src[pixG], src[pixB], src[pixA], src[pixL])
	//

	var i1 = math.MaxUint8 - uint32(src[pixA])

	//
	dst[pixR] = uint8(uint32(dst[pixR])*i1/math.MaxUint8 + uint32(src[pixR]))
	dst[pixG] = uint8(uint32(dst[pixG])*i1/math.MaxUint8 + uint32(src[pixG]))
	dst[pixB] = uint8(uint32(dst[pixB])*i1/math.MaxUint8 + uint32(src[pixB]))
	dst[pixA] = uint8(uint32(dst[pixA])*i1/math.MaxUint8 + uint32(src[pixA]))
}
// private usage
func mixRGBA32Src(dst []uint8, src ...uint8) {
	src[pixR], src[pixG], src[pixB], src[pixA] = intense(src[pixR], src[pixG], src[pixB], src[pixA], src[pixL])
	//

	dst[pixR] = src[pixR]
	dst[pixG] = src[pixG]
	dst[pixB] = src[pixB]
	dst[pixA] = src[pixA]
}
