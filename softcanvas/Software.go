package softcanvas

import (
	"github.com/iamGreedy/canvas"
	"github.com/iamGreedy/commons/colors"
	"github.com/pkg/errors"
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
	isclose int32
	//
	size      image.Point
	pix       []uint8
	mix       canvas.MixOperation
	precision Precision
	//
	pool *claaManager
}

func NewSoftware(w, h int, options ...canvas.Option) *Software {
	res := new(Software)
	res.size = image.Pt(w, h)
	res.precision = X16
	res.pool = newClaaManager(res)
	res.setup()
	for _, o := range options {
		if o.OptionType()&canvas.Init == canvas.Init {
			res.Option(o)
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
	case canvas.MixOperation:
		return true
	case *canvas.MixOperation:
		return true
	case Precision:
		return true
	case *Precision:
		return true
	}
	return false
}

// Surface interface
func (s *Software) Option(opt canvas.Option) canvas.Option {
	if atomic.CompareAndSwapInt32(&s.isclose, 1, 1) {
		return canvas.ResultFail{
			Cause: canvas.ErrorClosed,
		}
	}
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
		return o
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
		if o.Pix == nil{
			o.Pix = make([]uint8, len(s.pix))
		}
		copy(o.Pix, s.pix)
		return o
	case canvas.MixOperation:
		s.mix = o
		return o
	case *canvas.MixOperation:
		if o == nil {
			o = new(canvas.MixOperation)
		}
		*o = s.mix
		return o
	case Precision:
		s.precision = o
		return Precision(s.precision)
	case *Precision:
		if o == nil {
			o = new(Precision)
		}
		*o = Precision(s.precision)
		return o
	}
	return canvas.ResultFail{
		Cause: canvas.ErrorUnsupported,
	}
}


// Surface interface
func (s *Software) Query(query *canvas.Path, shader canvas.Shader, transform *canvas.Transform) error {
	if atomic.CompareAndSwapInt32(&s.isclose, 1, 1) {
		return canvas.ErrorClosed
	}
	if query == nil {
		return errors.WithMessage(canvas.ErrorNotAllowNil, "query can't be nil")
	}
	if shader == nil {
		shader = canvas.NewColorShader(colors.HTML.Black)
	}
	if transform == nil {
		transform = canvas.NewTransform()
	}
	query.RectValidate(s.size.X, s.size.Y)
	ws := s.pool.get(query.Rect)
	defer s.pool.put(ws)

	if *transform != *canvas.NewTransform() {
		for i, v := range query.Data {
			query.Data[i] = transform.RawMul(v)
		}
	}
	ws.data(query)
	//
	switch shd := shader.(type) {
	case *canvas.ColorShader:
		s.qColor(ws, shd, query)
	}

	return nil
}

func (s *Software) qColor(clb *claa, shd *canvas.ColorShader, qry *canvas.Path) error {
	var wg sync.WaitGroup
	wg.Add(qry.Rect.Max.Y - qry.Rect.Min.Y)
	if s.mix == canvas.Src {
		for y := qry.Rect.Min.Y; y < qry.Rect.Max.Y; y++ {
			go func(y int) {
				var prev int32 = 0
				for x := qry.Rect.Min.X; x < qry.Rect.Max.X; x++ {
					b, i := clb.offset(x, y)
					cell := clb.buffer[b][i]
					res := clb.effective(cell.cover+prev, cell.area)
					if cell.area != 0 || cell.cover != 0 {
						prev += cell.cover
					}
					if res == 0 {
						continue
					}
					offset := (s.size.X*y + x)*4
					s.pix[offset+pixR] = clampu8(uint32(shd.R) * res / math.MaxUint8)
					s.pix[offset+pixG] = clampu8(uint32(shd.G) * res / math.MaxUint8)
					s.pix[offset+pixB] = clampu8(uint32(shd.B) * res / math.MaxUint8)
					s.pix[offset+pixA] = clampu8(uint32(shd.A) * res / math.MaxUint8)
				}
				wg.Done()
			}(y)
		}
	} else {
		for y := qry.Rect.Min.Y; y < qry.Rect.Max.Y; y++ {
			go func(y int) {
				var prev int32 = 0
				for x := qry.Rect.Min.X; x < qry.Rect.Max.X; x++ {
					b, i := clb.offset(x, y)
					cell := clb.buffer[b][i]
					res := clb.effective(cell.cover+prev, cell.area)
					if cell.area != 0 || cell.cover != 0 {
						prev += cell.cover
					}
					if res == 0 {
						continue
					}
					srca := uint32(shd.A) * res / math.MaxUint8
					dsta := math.MaxUint8 - srca

					offset := (s.size.X*y + x) * pixL
					s.pix[offset+pixR] = clampu8((uint32(s.pix[offset+pixR])*dsta + uint32(shd.R)*srca)/math.MaxUint8)
					s.pix[offset+pixG] = clampu8((uint32(s.pix[offset+pixG])*dsta + uint32(shd.G)*srca)/math.MaxUint8)
					s.pix[offset+pixB] = clampu8((uint32(s.pix[offset+pixB])*dsta + uint32(shd.B)*srca)/math.MaxUint8)
					s.pix[offset+pixA] = clampu8((uint32(s.pix[offset+pixA])*dsta + uint32(shd.A)*srca)/math.MaxUint8)
				}
				wg.Done()
			}(y)
		}
	}
	wg.Wait()
	return nil
}

// Surface interface
func (s *Software) Close() error {
	atomic.StoreInt32(&s.isclose, 1)
	return nil
}

// Convenient Method
func (s *Software) Clear(c color.RGBA) {
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
}

// private usage
func clampu8(a uint32) uint8 {
	if a > math.MaxUint8 {
		return math.MaxUint8
	}
	return uint8(a)
}
