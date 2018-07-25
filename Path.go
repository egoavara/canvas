package canvas

import (
	"github.com/go-gl/mathgl/mgl32"
	"math"
	"io"
	"github.com/iamGreedy/commons/align"
	"image"
)

const capData = 128
const (
	SVGQuery QueryType = iota
)
var Spacer = mgl32.Vec3{float32(math.NaN()), float32(math.NaN()), float32(math.NaN())}

type (
	Path struct {
		Data []mgl32.Vec3
		Rect image.Rectangle
	}
	InnerPath interface {
		MoveTo(to mgl32.Vec2)
		LineTo(to mgl32.Vec2)
		QuadTo(p0, to mgl32.Vec2)
		CubeTo(p0, p1, to mgl32.Vec2)
		CloseTo()

		Query(qtype QueryType, reader io.Reader) error
	}
	TextPath interface {
		Follow(text string, pathfn func(i InnerPath))
		Draw(text string, pos mgl32.Vec2, align align.Align)
	}
	strokeInnerPath struct {
		to *Path
		pen  mgl32.Vec2
		from mgl32.Vec2
	}
	QueryType uint32
)

func NewPath() *Path {
	return &Path{
		Data: make([]mgl32.Vec3, 0, capData),
	}
}
func (s *Path ) RectValidate(w, h int) {
	s.Rect.Min.X = iMax(s.Rect.Min.X, 0)
	s.Rect.Min.Y = iMax(s.Rect.Min.Y, 0)
	s.Rect.Max.X = iMin(s.Rect.Max.X, w)
	s.Rect.Max.Y = iMin(s.Rect.Max.Y, h)
}

func (s *Path ) Fill(fn func(i InnerPath)) *Path{
	res := &fillInnerPath{
		to:s,
	}
	fn(res)
	s.Rect.Min.X = int(res.min[0])
	s.Rect.Min.Y = int(res.min[1])
	s.Rect.Max.X = int(res.max[0])
	s.Rect.Max.Y = int(res.max[1])
	return s
}

// TODO : func (s *Path) Stroke(fn func(i InnerPath)) {}
// TODO : func (s *Path) FillText(fn func(t TextPath)) {}
// TODO : func (s *Path) StrokeText(fn func(t TextPath)) { }

func (s *Path) SetFont(f *Font) {

}
func (s *Path) SetLineWidth(width float32) {

}
func (s *Path) SetLineJoint(width float32) {

}
func (s *Path) SetLineCap(width float32) {

}
func iMin(a, b  int) int {
	if a < b{
		return a
	}
	return b
}
func iMax(a, b  int) int {
	if a > b{
		return a
	}
	return b
}