package canvas

import "io"

type Surface interface {
	// Check Support options
	// Returns true if all the given options are satisfied
	Support(opt ... Option) bool
	// It is highly affected by the Linux ioctl function.
	// You can pass Option pointer or value(pointer is normally use for get, value normally used for set)
	//
	// return can be error
	// If there is error, it return nil or ResultFail struct(not pointer of ResultFail)
	Option(opt Option) Option
	// Query send Path, Shader, Transform for Path Rendering
	// Always it MUST BE thread-safe
	// You can use Goroutine
	//
	// throw > ErrorTooBusy : Literally, it is too busy to handle this
	// throw > ErrorClosed 	: .Close() method called, do not use this Surface
	// throw > * : any other errors
	//
	// query 		: It is points for path render
	//				:
	//				: nil not allowed, you must pass valid path data
	//				throw > ErrorNotAllowNil
	//
	// shader		: It is filling data for rendering area
	// 				: This will determine how to fill given space
	//				:
	//				: nil can be use, it MUST parsed as NewColorShader(colors.HTML.Black)
	//				throw > _
	//
	// transform 	: It is 3x3 matrix for homogeneous vector points(in query)
	//				: Every points in query is handle by this for transfomation
	//				: You can use directly apply this by handling Path.Data
	//				: But query build is CPU-dependent work
	//				: Surface can be accelerate by GPU
	//				: So if you can, pass transformation by using this
	//				:
	//				: nil can be use, it assume Identity matrix
	//				throw > _
	Query(query *Path, shader Shader, transform *Transform) error
	// This make free every resource for surface
	// It destroy Surface
	// Even if Surface call this, it is possible to use Surface in some case(not using external resource like GPU)
	// But not recommanded for safety reason
	io.Closer
}