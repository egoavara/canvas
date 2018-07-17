package canvas

import (
	"image"
	"image/draw"
)

type Surface interface {
	// Return SurfaceType
	// This tells you how the surface will be rendered
	// For more infomation see SurfaceType comment
	//
	Type() SurfaceType

	// Get or set function for surface
	// It can be many option can pass
	// First arguments, 'opt' is choose this is getter or setter
	// Care to set this value with goroutine, this is not thread-safe
	//
	// * getter
	// If 'opt' is Pointer(ex:'*AnyType'), it is getter
	// If getter, it return NONE POINTER VALUE that specified type
	// But also set option result to given pointer like C pointer argument output
	// 		ex) var size = canvas.Surface(v).Option(new(canvas.OptionSize)).(canvas.OptionSize)
	//		ex) var size canvas.OptionSize
	//			canvas.Surface(v).Option(&size)
	//
	// * setter
	// If 'opt' is NonePointer(ex:'AnyType'), it is setter
	// If setter, it return changed value with NONE POINTER
	//		ex) Surface(v).Option(canvas.OptionSize{X:1024, Y:1024})
	Option(opt interface{}) interface{}

	// Request query
	// Always it is thread-safe
	// You can call query with goroutine
	//
	// - thread-safe
	//
	// > query 		: It is points for path render
	//				  - not nilable
	// > shader		: It is filling data for rendering area
	//			  	  It decide how to paint given area
	//			  	  It can be nil, when this is nil, it is assumed NewColorShader(colors.HTML.Black)
	//				  - nilable
	// > transform 	: It is 4x4 transformation matrix for each point(in 'query')
	//				  It can be nil, when this is nil, it is assumed NewTransform() == identity matrix 4x4
	//				  - nilable
	// throw
	// 		: ErrorTooBusy		:
	// 		: ErrorUnsupported	:
	// 		: ErrorNoQuery		:
	// 		: ErrorInvalidData	:
	Query(query *Path, shader Shader, transform *Transform) error

	// Cancel all requested query
	// Remove all non-flush queries
	Cancel()
	// Apply all queries to dst image
	Flush() error
	// image clear with color RGBA(0, 0, 0, 0)
	Clear() error
	Draw(dst draw.Image, r image.Rectangle, sp image.Point)
}