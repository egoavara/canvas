package canvas

import (
	"math"
	"github.com/go-gl/mathgl/mgl32"
)

var MinusZero = math.Float32frombits(0x80000000)

func floatSign(a float32, sign uint32) float32 {
	return math.Float32frombits(math.Float32bits(a) ^ sign)
}
func sign(dir int, a float32) int {
	dir += int(math.Float32bits(a) >> 31)
	if dir < -1{
		return -1
	}
	if dir > 1 {
		return 1
	}
	return 0
}

func devSquared(a, b, c mgl32.Vec2) float32 {
	devx := a[0] - 2*b[0] + c[0]
	devy := a[1] - 2*b[1] + c[1]
	return devx*devx + devy*devy
}


func lerp(t float32, p, q mgl32.Vec2) mgl32.Vec2 {
	return [2]float32{p[0] + t*(q[0]-p[0]), p[1] + t*(q[1]-p[1])}
}