package softcanvas

import (
	"github.com/iamGreedy/canvas"
	"github.com/iamGreedy/canvas/claa"
	"image"
	"image/color"
	"image/draw"
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
	size image.Point
	pix  []uint8
	mix  canvas.MixOperation
	//
	claa *claa.CLAA
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

	//	switch o := opt.(type) {
	//	case OptionCLAAPrecision:
	//	case *OptionCLAAPrecision:
	//	case canvas.Width:
	//		s.size.X = int(o)
	//		s.setup(s.claa.GetPrecision())
	//		return canvas.Width(s.size.X)
	//	case *canvas.Width:
	//		if o == nil{
	//			o = new(canvas.Width)
	//		}
	//		*o = canvas.Width(s.size.X)
	//	case canvas.Height:
	//		return canvas.Height(s.size.Y)
	//	case *canvas.Height:
	//		if o == nil{
	//			o = new(canvas.Height)
	//		}
	//		*o = canvas.Height(s.size.Y)
	//		return *o
	//	case canvas.Size:
	//		return canvas.Size(s.size)
	//	case *canvas.Size:
	//		if o == nil{
	//			o = new(canvas.Size)
	//		}
	//		*o = canvas.Size(s.size)
	//		return *o
	//	}
	//	return nil

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
	case *canvas.Height:
	case canvas.Size:
	case *canvas.Size:
	case *canvas.Driver:
	case canvas.Clear:
	case canvas.FrameBuffer:
	case *canvas.FrameBuffer:
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
		s.claa.SetPrecision(claa.Precision(o))
		return OptionCLAAPrecision(s.claa.GetPrecision())
	case *OptionCLAAPrecision:
		if o == nil {
			o = new(OptionCLAAPrecision)
		}
		*o = OptionCLAAPrecision(s.claa.GetPrecision())
		return *o
	}
	return nil
}

func (s *Software) Query(query *canvas.Path, shader canvas.Shader, transform *canvas.Transform) error {
	panic("implement me")
}

func NewSoftware(w, h int, options ...canvas.Option) *Software {
	res := new(Software)
	res.size = image.Pt(w, h)
	res.pix = make([]uint8, w*h*pixL)
	res.claa = claa.NewCLAA(claa.X16, w, h)
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

func (s *Software) setup() {
	s.pix = make([]uint8, s.size.X*s.size.Y*pixL)
	s.claa = claa.NewCLAA(s.claa.GetPrecision(), s.size.X, s.size.Y)
}
