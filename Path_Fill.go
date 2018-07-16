package canvas

import (
	"github.com/go-gl/mathgl/mgl32"
	"io"
	"math"
)

type fillInnerPath struct {
	to   *Path
	pen  mgl32.Vec2
	from mgl32.Vec2
}

func (s *fillInnerPath) MoveTo(to mgl32.Vec2) {
	s.to.Data = append(s.to.Data, to.Vec3(1))
	s.from = to
	s.pen = to
}
func (s *fillInnerPath) LineTo(to mgl32.Vec2) {
	s.to.Data = append(s.to.Data, to.Vec3(1))
	s.pen = to

}
func (s *fillInnerPath) QuadTo(p0, to mgl32.Vec2) {
	from := s.pen
	devsq := devSquared(from, p0, to)
	if devsq >= 0.333 {
		const tol = 3
		n := 1 + int(math.Sqrt(math.Sqrt(tol*float64(devsq))))
		t, nInv := float32(0), 1/float32(n)
		for i := 0; i < n-1; i++ {
			t += nInv
			fromPivot := lerp(t, from, p0)
			pivotTo := lerp(t, p0, to)
			s.LineTo(lerp(t, fromPivot, pivotTo))
		}
	}
	s.LineTo(to)
}
func (s *fillInnerPath) CubeTo(p0, p1, to mgl32.Vec2) {
	from := s.pen
	devsq := devSquared(from, p0, to)
	if devsqAlt := devSquared(from, p1, to); devsq < devsqAlt {
		devsq = devsqAlt
	}
	if devsq >= 0.333 {
		const tol = 3
		n := 1 + int(math.Sqrt(math.Sqrt(tol*float64(devsq))))
		t, nInv := float32(0), 1/float32(n)
		for i := 0; i < n-1; i++ {
			t += nInv
			ab := lerp(t, from, p0)
			bc := lerp(t, p0, p1)
			cd := lerp(t, p1, to)
			abc := lerp(t, ab, bc)
			bcd := lerp(t, bc, cd)
			s.LineTo(lerp(t, abc, bcd))
		}
	}
	s.LineTo(to)
}
func (s *fillInnerPath) CloseTo() {
	s.to.Data = append(s.to.Data, s.from.Vec3(1), Spacer)
}

func (s *fillInnerPath) Query(qtype QueryType, reader io.Reader) {
	s.to.Data = append(s.to.Data, s.from.Vec3(1), Spacer)
}
