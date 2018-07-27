package canvas

import (
	"github.com/go-gl/mathgl/mgl32"
)

func devSquared(a, b, c mgl32.Vec2) float32 {
	devx := a[0] - 2*b[0] + c[0]
	devy := a[1] - 2*b[1] + c[1]
	return devx*devx + devy*devy
}


func lerp(t float32, p, q mgl32.Vec2) mgl32.Vec2 {
	return [2]float32{p[0] + t*(q[0]-p[0]), p[1] + t*(q[1]-p[1])}
}

func Vec(x, y float32) mgl32.Vec2 {
	return mgl32.Vec2{
		x, y,
	}
}