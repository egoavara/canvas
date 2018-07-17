package softcanvas

import (
	"image"
	"image/draw"
	"github.com/iamGreedy/canvas/claa"
	"github.com/iamGreedy/canvas"
)

const (
	r, g, b, a = 0, 1, 2, 3
)

type Software struct {
	//
	size image.Point
	pix  []uint8
	//
	claa *claa.CLAA
}

func NewSoftware(w, h int, options ... interface{}) *Software {
	res := new(Software)
	res.size = image.Pt(w, h)
	res.claa = claa.NewCLAA(claa.X16, w, h)
	for _, opt := range options {
		res.Option(opt)
	}
	return res
}

func (s *Software) Option(opt interface{}) interface{} {
	switch o := opt.(type) {
	case OptionCLAAPrecision:
		s.claa.SetPrecision(claa.Precision(o))
		return OptionCLAAPrecision(s.claa.GetPrecision())
	case *OptionCLAAPrecision:
		if o == nil{
			o = new(OptionCLAAPrecision)
		}
		*o = OptionCLAAPrecision(s.claa.GetPrecision())
		return *o
	case canvas.OptionWidth:
		s.size.X = int(o)
		s.setup(s.claa.GetPrecision())
		return canvas.OptionWidth(s.size.X)
	case *canvas.OptionWidth:
		if o == nil{
			o = new(canvas.OptionWidth)
		}
		*o = canvas.OptionWidth(s.size.X)
	case canvas.OptionHeight:
		return canvas.OptionHeight(s.size.Y)
	case *canvas.OptionHeight:
		if o == nil{
			o = new(canvas.OptionHeight)
		}
		*o = canvas.OptionHeight(s.size.Y)
		return *o
	case canvas.OptionSize:
		return canvas.OptionSize(s.size)
	case *canvas.OptionSize:
		if o == nil{
			o = new(canvas.OptionSize)
		}
		*o = canvas.OptionSize(s.size)
		return *o
	}
	return nil
}

func (s *Software) Type() canvas.SurfaceType {
	return canvas.SurfaceTypeSoftware
}


func (s *Software) Clear() error {
	s.claa.Clear()


}

func (s *Software) Query(query *canvas.Path, shader canvas.Shader, transform *canvas.Transform) error {
	panic("implement me")
}

func (s *Software) Draw(dst draw.Image, r image.Rectangle, sp image.Point) {
	panic("implement me")
}

func (s *Software ) setup(p claa.Precision) {
	s.pix = make([]uint8, s.size.X * s.size.Y * 4)
	s.claa = claa.NewCLAA(p, s.size.X, s.size.Y)
}