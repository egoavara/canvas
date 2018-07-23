package canvas

import (
	"image"
	"image/color"
	"image/draw"
)

type Option interface {
	OptionType() OptionType
}
type OptionType uint8

// When you use Option with pointer, it is Get(call getter)
// If it is non-pointer, it is Set(call setter)
//
// If Option is setter .Option() return nil or error(when fail to set)
// If Option is getter .Option() retunr Pointer of passed value, or nil(if it was not read available Option)
//
const (
	Read  OptionType = 0x01
	Write OptionType = 0x02
	Init  OptionType = 0x04
	Send  OptionType = 0x08
	Error OptionType = 0x80
	//
	// Read  		OptionType
	// Write  		OptionType
	// Init  		OptionType
	// Send  		OptionType
	//
	// empty  		OptionType
	// empty 		OptionType
	// empty 		OptionType
	// ErrorFlag  	OptionType
)

// There is two type Fail, one is .Option(...) return error wrap by ResultFail
// Other is return nil, which is so obious, no need to explain
//
type ResultFail struct {
	msg error
}
func (ResultFail) OptionType() OptionType {
	return Error
}
func (s ResultFail) Error() string {
	return s.msg.Error()
}
func IsFail(o Option) bool {
	if o == nil{
		return true
	}
	return o.OptionType() & Error == Error
}

// Essensial support option
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
func (s FrameBuffer) Convert() *image.RGBA {
	return &image.RGBA{
		Stride: s.Stride,
		Rect:   s.Rect,
		Pix:    s.Pix,
	}
}
func (MixOperation) OptionType() OptionType {
	return Read | Write | Init
}

var (
	Src  = MixOperation(draw.Src)
	Over = MixOperation(draw.Over)
)

// Buffer support option
type (
	// Use Flush buffer for performance
	//
	// = Send
	//
	UseBuffer struct{}

	// GetCurrent Buffer Query Count
	//
	// = Read
	//
	BufferLength int

	// Close buffer
	//
	// = Send
	CloseBuffer struct{}

	// Flush all buffer for Data
	//
	// = Send
	FlushBuffer struct{}

	// Lock until every flushing success
	//
	// = Send
	WaitFlush struct{}
)

func (UseBuffer) OptionType() OptionType {
	return Send
}
func (BufferLength) OptionType() OptionType {
	return Read
}
func (CloseBuffer) OptionType() OptionType {
	return Send
}
func (FlushBuffer) OptionType() OptionType {
	return Send
}
func (WaitFlush) OptionType() OptionType {
	return Send
}
