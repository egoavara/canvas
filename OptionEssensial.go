package canvas

import (
	"image"
	"image/color"
	"image/draw"
)

var checkvaliddriver = []Option{
	Width(1),
	new(Width),
	Height(1),
	new(Height),
	Size{},
	new(Size),
	Clear{255,255,255,255},
	FrameBuffer{
		Pix: []uint8{255,255,255,255},
		Rect:image.Rect(0,0,1,1),
		Stride:4,
	},
	new(FrameBuffer),
	MixOperation(Src),
	new(MixOperation),
}
// Essensial support option
// These Option MUST implement
type (
	// Get Width
	//
	// = Read | Write
	Width int
	// Get Height
	//
	// = Read | Write
	Height int
	// Get Size(Width, Height)
	//
	// = Read | Write
	Size image.Point
	// Get DriverName
	//
	// = Read
	Driver string
	// Clear with color.RGBA
	// Clear{} for using this empty all pixels
	//
	// = Write | Init
	//
	Clear color.RGBA
	// GetResult image
	//
	// = Read | Write
	FrameBuffer image.RGBA
	// GetResult image
	//
	// = Read | Write | Init
	MixOperation draw.Op
)


func (Width) OptionType() OptionType {
	return Read | Write
}
func (Height) OptionType() OptionType {
	return Read | Write
}
func (Size) OptionType() OptionType {
	return Read | Write
}
func (Driver) OptionType() OptionType {
	return Read
}
func (Clear) OptionType() OptionType {
	return Write | Init
}
func (FrameBuffer) OptionType() OptionType {
	return Read | Write
}
func (MixOperation) OptionType() OptionType {
	return Read | Write | Init
}
func (s FrameBuffer) Convert() *image.RGBA {
	return &image.RGBA{
		Stride: s.Stride,
		Rect:   s.Rect,
		Pix:    s.Pix,
	}
}
var (
	Src  = MixOperation(draw.Src)
	Over = MixOperation(draw.Over)
)