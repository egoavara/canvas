package canvas

import (
	"github.com/iamGreedy/commons/colors"
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
	// Set value to Surface
	// getter
	Read OptionType = 1 << iota
	// Set value to Surface
	// setter
	// ex) .Option(<Write>(value))
	Write OptionType = 1 << iota
	// Init is special case for write which can pass to constructor
	// setter
	// ex) NewSurface(w, h, <Init>)
	Init OptionType = 1 << iota
	Send OptionType = 1 << iota
	// Can be delay syncro reason
	Sync OptionType = 1 << iota
	//
	// reserve0  	OptionType = 1 << iota
	// reserve1 	OptionType = 1 << iota
	// empty 		OptionType = 1 << iota
	// empty  		OptionType = 1 << iota
)

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
func (MixOperation) OptionType() OptionType {
	return Read | Write| Init
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
