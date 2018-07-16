package canvas

import "math"

const (
	SurfaceTypeSoftware SurfaceType = iota
	SurfaceTypeOpenGL   SurfaceType = iota
	// TODO : OpenCL SurfaceType = iota
	SurfaceTypeInvalid = math.MaxUint32
)

type SurfaceType uint32

var _SurfaceType = map[SurfaceType]string{
	SurfaceTypeSoftware: "Legacy",
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
