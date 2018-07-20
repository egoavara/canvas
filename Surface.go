package canvas

import (
	"image"
	"image/draw"
)

type Surface interface {
	// Check Support options
	Support(opt ... Option) bool
	// Get or set function for surface
	// It can be many option can pass
	// First arguments, 'opt' is choose this is getter or setter
	// Care to set this value with goroutine, this is not thread-safe
	//
	// * getter
	// If 'opt' is Pointer(ex:'*AnyType'), it is getter
	// If getter, it return NONE POINTER VALUE that specified type
	// But also set option result to given pointer like C pointer argument output
	// 		ex) var size = canvas.Surface(v).Option(new(canvas.Size)).(canvas.Size)
	//		ex) var size canvas.Size
	//			canvas.Surface(v).Option(&size)
	//
	// * setter
	// If 'opt' is NonePointer(ex:'AnyType'), it is setter
	// If setter, it return changed value with NONE POINTER
	//		ex) Surface(v).Option(canvas.Size{X:1024, Y:1024})
	Option(opt Option) Option
	// Request query
	// Always it is thread-safe
	// You can call query with goroutine
	//
	// - thread-safe, you can call this on goroutine
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
}