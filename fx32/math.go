package fx32

import (
	"golang.org/x/image/math/fixed"
	"github.com/go-gl/mathgl/mgl32"
	"strconv"
	"fmt"
)

const fixed32converter = 1 << 6
const precision = 0x3F
const exponent = ^precision
const (
	Quarter = 0x10
	Half    = 0x20
	One     = 0x40
	MinusOne = -0x40
)

func F32ToFx32(x float32) fixed.Int26_6 {
	return fixed.Int26_6(x * fixed32converter)
}
func Fx32ToF32(x fixed.Int26_6) float32 {
	return float32(x) / fixed32converter
}

func V32ToVx32(x mgl32.Vec2) fixed.Point26_6 {
	return fixed.Point26_6{
		X:fixed.Int26_6(x[0] * fixed32converter) >> 4 << 4,
		Y:fixed.Int26_6(x[1] * fixed32converter) >> 4 << 4,
	}
}
func Vx32ToV32(x fixed.Point26_6) mgl32.Vec2{
	return mgl32.Vec2{
		float32(x.X) / fixed32converter,
		float32(x.Y) / fixed32converter,
	}
}

func F32ToString(x fixed.Int26_6) string {
	return strconv.FormatFloat(float64(Fx32ToF32(x)), 'f', -1, 32)
}

func V32ToString(x fixed.Point26_6) string {
	return fmt.Sprintf("Vx32{X: %s, Y:%s}", F32ToString(x.X), F32ToString(x.Y))
}

func Abs(int26_6 fixed.Int26_6) fixed.Int26_6 {
	if int26_6 < 0 {
		return -int26_6
	}
	return int26_6
}
func Clamp(int26_6, min, max fixed.Int26_6) fixed.Int26_6 {
	if int26_6 < min {
		return min
	}
	if int26_6 > max {
		return max
	}
	return int26_6
}
func Ceil(int26_6 fixed.Int26_6) fixed.Int26_6 {
	if int26_6&precision == 0 {
		return int26_6 & exponent
	}
	return int26_6&exponent + One
}
func Round(int26_6 fixed.Int26_6) fixed.Int26_6 {
	return (int26_6 + Half) & exponent
}
func Floor(int26_6 fixed.Int26_6) fixed.Int26_6 {
	return int26_6 & exponent
}
