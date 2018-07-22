package softcanvas

import (
	"github.com/iamGreedy/canvas"
	"github.com/iamGreedy/canvas/claa"
	"image"
	"image/color"
	"image/draw"
	"sync"
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
	//
	linelck []*sync.Mutex
	works   *sync.WaitGroup
	//
	q map[canvas.Shader]q
	//
	pool *sync.Pool
}
type q struct {
	p *canvas.Path
	t *canvas.Transform
}

func NewSoftware(w, h int, options ...canvas.Option) *Software {
	res := new(Software)
	res.size = image.Pt(w, h)
	res.precision = claa.X16
	res.works = new(sync.WaitGroup)
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

func (s *Software) Support(opt ...canvas.Option) bool {
	for _, o := range opt {
		if !s.support(o) {
			return false
		}
	}
	return true
}
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

	case *canvas.UseBuffer:
	case *canvas.BufferLength:
	case canvas.CloseBuffer:
	case *canvas.CloseBuffer:
	case canvas.FlushBuffer:
	case *canvas.FlushBuffer:
	case canvas.WaitFlush:
	case *canvas.WaitFlush:
		//
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

func (s *Software) Query(query *canvas.Path, shader canvas.Shader, transform *canvas.Transform) error {
	panic("implement me")
}
func (s *Software) query(query *canvas.Path, shader canvas.Shader, transform *canvas.Transform) {
	go func() {

	}()
}

// private
func (s *Software) setup() {
	s.pix = make([]uint8, s.size.X*s.size.Y*pixL)
	s.linelck = make([]*sync.Mutex, s.size.Y)
	s.pool = &sync.Pool{
		New: func() interface{} {
			return claa.NewCLAA(s.precision, s.size.X, s.size.Y)
		},
	}
	for i := range s.linelck {
		s.linelck[i] = new(sync.Mutex)
	}
	//s.claa = claa.NewCLAA(s.claa.GetPrecision(), s.size.X, s.size.Y)
}
func (s *Software) usebuf() {
	if s.q == nil {
		s.q = make(map[canvas.Shader]q)
	}
}
func (s *Software) flush() {
	for shader, q := range s.q {
		go s.query(q.p, shader, q.t)
	}
}
func (s *Software) closebuf() {
	s.q = nil
}

//
func (s *Software) Clear(c color.RGBA) {
	for y := 0; y < s.size.Y; y++ {
		for x := 0; x < s.size.X; x++ {
			offset := s.size.X*y + x*pixL
			s.pix[offset+pixR] = c.R
			s.pix[offset+pixG] = c.G
			s.pix[offset+pixB] = c.B
			s.pix[offset+pixA] = c.A
		}
	}
}
