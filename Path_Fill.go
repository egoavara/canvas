package canvas

import (
	"github.com/go-gl/mathgl/mgl32"
	"io"
	"math"
	"github.com/iamGreedy/psvg"
	"github.com/pkg/errors"
	"fmt"
)

type fillInnerPath struct {
	to   *Path
	pen  mgl32.Vec2
	from mgl32.Vec2
	//
	min mgl32.Vec2
	max mgl32.Vec2
}

func (s *fillInnerPath) MoveTo(to mgl32.Vec2) {
	s.to.Data = append(s.to.Data, to.Vec3(1))
	s.from = to
	s.pen = to
	//
	s.min[0] = f32min(s.min[0], to[0])
	s.min[1] = f32min(s.min[1], to[1])
	s.max[0] = f32max(s.max[0], to[0])
	s.max[1] = f32max(s.max[1], to[1])
}
func (s *fillInnerPath) LineTo(to mgl32.Vec2) {
	s.to.Data = append(s.to.Data, to.Vec3(1))
	s.pen = to
	//
	s.min[0] = f32min(s.min[0], to[0])
	s.min[1] = f32min(s.min[1], to[1])
	s.max[0] = f32max(s.max[0], to[0])
	s.max[1] = f32max(s.max[1], to[1])
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

func (s *fillInnerPath) Query(qtype QueryType, reader io.Reader) error {
	//s.to.Data = append(s.to.Data, s.from.Vec3(1), Spacer)
	switch qtype {
	case SVGQuery:
		p, err := psvg.NewRendererFromReader(reader)
		if err != nil {
			return err
		}
		p.Render(s)
	}
	return errors.New(fmt.Sprintf("Unknown QueryType(%v)", qtype))
}

func f32min(a, b float32) float32 {
	return float32(math.Min(float64(a), float64(b)))
}
func f32max(a, b float32) float32 {
	return float32(math.Max(float64(a), float64(b)))
}