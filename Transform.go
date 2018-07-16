package canvas

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/iamGreedy/axis"
)

type Transform mgl32.Mat3

func NewTransform() *Transform {
	return &Transform{1, 0, 0, 0, 1, 0, 0, 0, 1}
}
//
func (s *Transform) Rotate(degree float32) *Transform {
	*s = Transform(mgl32.HomogRotate2D(mgl32.DegToRad(degree)).Mul3(mgl32.Mat3(*s)))
	return s
}
func (s *Transform) Translate(delta mgl32.Vec2) *Transform {
	*s = Transform(mgl32.Translate2D(delta[0], delta[1]).Mul3(mgl32.Mat3(*s)))
	return s
}
func (s *Transform) Scale(scale mgl32.Vec2) *Transform {
	*s = Transform(mgl32.Scale2D(scale[0], scale[1]).Mul3(mgl32.Mat3(*s)))
	return s
}
func (s *Transform) Reflection(a axis.Axis) *Transform {
	var v, h float32 = 1, 1
	if a.HasVertical() {
		v = -1
	}
	if a.HasHorizontal() {
		h = -1
	}
	*s = Transform(mgl32.Mat3{
		v, 0, 0,
		0, h, 0,
		0, 0, 1,
	}.Mul3(mgl32.Mat3(*s)))
	return s
}
func (s *Transform) Shear(shear float32, ea axis.ExculsiveAxis) *Transform {
	switch ea.Axis() {
	case axis.Vertical:
		*s = Transform(mgl32.ShearX2D(shear).Mul3(mgl32.Mat3(*s)))
	case axis.Horizontal:
		*s = Transform(mgl32.ShearY2D(shear).Mul3(mgl32.Mat3(*s)))
	}
	return s
}
func (s *Transform) Pivot(at mgl32.Vec2, do func(t *Transform)) *Transform {
	s.Translate(at.Mul(-1))
	do(s)
	return s.Translate(at)
}

//
func (s *Transform) Mul(v mgl32.Vec2) mgl32.Vec2 {
	return s.rawMul(v.Vec3(1)).Vec2()
}
func (s *Transform) rawMul(v mgl32.Vec3) mgl32.Vec3 {
	return mgl32.Mat3(*s).Mul3x1(v)
}
func (s *Transform) String() string {
	return mgl32.Mat3(*s).String()
}
