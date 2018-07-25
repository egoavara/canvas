package canvas

import (
	"image"
	"image/color"
)

type Shader interface {
	Size(w, h int)
	At(x, y int) color.RGBA
}
type (
	ColorShader color.RGBA
	FixedShader struct {
		data *image.RGBA
	}
	RepeatShader struct {
		data *image.RGBA
	}
	KernalShader struct {
		data *image.RGBA
	}
	//LinearGradientShader struct {}
	//RadialGradientShader struct {}
)

func NewColorShader(c color.Color) Shader {
	cc := color.RGBAModel.Convert(c).(color.RGBA)
	return &ColorShader{R: cc.R, G: cc.G, B: cc.B, A: cc.A}
}
func (s *ColorShader) Size(w, h int) {}
func (s *ColorShader) At(x, y int) color.RGBA {
	return color.RGBA(*s)
}

func (s *FixedShader) Size(w, h int) {

}
func (s *FixedShader) At(x, y int) color.RGBA {
	offset := s.data.Stride*y + x*4
	return color.RGBA{
		R: s.data.Pix[offset+0],
		G: s.data.Pix[offset+1],
		B: s.data.Pix[offset+2],
		A: s.data.Pix[offset+3],
	}
}
