package canvas

import "math"

const (
	SurfaceTypeSoftware SurfaceType = iota
	SurfaceTypeOpenGL   SurfaceType = iota
	// TODO : OpenCL SurfaceType = iota
	// TODO : Skia SurfaceType = iota
	// TODO : SDL2 SurfaceType = iota
	SurfaceTypeInvalid = math.MaxUint32
)

// The SurfaceType tells you which rendering tool is used to draw the surface.
// But there is many different way to render
// So using same SurfaceType those not mean they are draw same way, Like algorithm, pipeline, ...
//
//
// - SurfaceTypeSoftware 	: Use CPU
// - SurfaceTypeOpenGL 		: Use OpenGL(https://www.opengl.org/)
type SurfaceType uint32
var _SurfaceType = map[SurfaceType]string{
	SurfaceTypeSoftware: "Software",
	SurfaceTypeOpenGL:   "OpenGL",
}
func ToSurfaceType(src string) SurfaceType {
	for key, value := range _SurfaceType {
		if value == src {
			return key
		}
	}

	return SurfaceTypeInvalid
}
func (s SurfaceType) String() string {
	v, ok := _SurfaceType[s]
	if ok {
		return v
	}
	return "< invalid >"
}
