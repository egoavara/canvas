package canvas

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/iamGreedy/commons/align"
)

// TODO : var Figure _Figure
type _Figure struct {}
func (_Figure) Circle(i InnerPath, center mgl32.Vec2, radius, startAngle, endAngle float32) {

}


func (_Figure) Ellipse(i InnerPath, center mgl32.Vec2, radius mgl32.Vec2, startAngle, endAngle float32) {

}
func (_Figure) RegularPolygon(i InnerPath, center mgl32.Vec2, n int, rotation float32) {

}
func (_Figure) Polygon(i InnerPath, points ... mgl32.Vec2) {

}

func (_Figure) Rect(i InnerPath, posalign align.Align, position mgl32.Vec2, size mgl32.Vec2) {

}
