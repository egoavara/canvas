package canvas

import (
	"image"
	"image/draw"
	"github.com/pkg/errors"
)

type Surface interface {
	Type() SurfaceType
	//
	Bounds() image.Rectangle
	Clear() error
	Query(query *Path, shader Shader, transform *Transform) error
	Draw(dst draw.Image, r image.Rectangle, sp image.Point)
}

func NewSurface(width, height int, tp SurfaceType) (Surface, error) {
	switch tp {
	case SurfaceTypeSoftware:
		//return &Legacy{
		//	buffer:     image.NewRGBA(image.Rect(0, 0, width, height)),
		//	swapbuffer: image.NewRGBA(image.Rect(0, 0, width, height)),
		//}, nil
		return &Software{
		}, nil
	case SurfaceTypeOpenGL:
		return nil, errors.New("Unavilable")
	default:
		return nil, errors.New("Unknown SurfaceType")
	}

}
