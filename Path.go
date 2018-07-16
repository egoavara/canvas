package canvas

import (
	"github.com/go-gl/mathgl/mgl32"
	"math"
	"io"
	"github.com/iamGreedy/align"
)

const capData = 128
const (
	SVGQuery QueryType = iota
)
var Spacer = mgl32.Vec3{float32(math.NaN()), float32(math.NaN()), float32(math.NaN())}

type (
	Path struct {
		Data []mgl32.Vec3
	}
	InnerPath interface {
		MoveTo(to mgl32.Vec2)
		LineTo(to mgl32.Vec2)
		QuadTo(p0, to mgl32.Vec2)
		CubeTo(p0, p1, to mgl32.Vec2)
		CloseTo()

		Query(qtype QueryType, reader io.Reader)
	}
	TextPath interface {
		Follow(text string, pathfn func(i InnerPath))
		Draw(text string, pos mgl32.Vec2, align align.Align)
	}
	QueryType uint32
	strokeInnerPath struct {
		to *Path
		pen  mgl32.Vec2
		from mgl32.Vec2
	}
)




func NewPath() *Path {
	return &Path{
		Data: make([]mgl32.Vec3, 0, capData),
	}
}

func (s *Path ) Fill(fn func(i InnerPath)) *Path{
	fn(&fillInnerPath{
		to:s,
	})
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
