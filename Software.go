package canvas

import (
	"image"
	"image/draw"
)

const (
	r, g, b, a = 0, 1, 2, 3
)

type Software struct {
	//
	size image.Point
	pix  []uint8
	//
	//buffer []cell
}

func (s *Software) Type() SurfaceType {
	panic("implement me")
}

func (s *Software) Bounds() image.Rectangle {
	panic("implement me")
}

func (s *Software) Clear() error {
	panic("implement me")
}

func (s *Software) Query(query *Path, shader Shader, transform *Transform) error {
	panic("implement me")
}

func (s *Software) Draw(dst draw.Image, r image.Rectangle, sp image.Point) {
	panic("implement me")
}


