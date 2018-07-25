package canvas_test

import (
	"fmt"
	"github.com/iamGreedy/canvas"
	"github.com/go-gl/mathgl/mgl32"
)

func ExampleNewTransform() {
	// Identity matrix
	fmt.Println(canvas.NewTransform())

	// translation matrix (+5, +5)
	fmt.Println(canvas.NewTransform().Translate(mgl32.Vec2{5,5}))

	// Scale (16x, 16x)
	fmt.Println(canvas.NewTransform().Scale(mgl32.Vec2{16,16}))

	// Scale and Translate
	fmt.Println(canvas.NewTransform().Scale(mgl32.Vec2{16,16}).Translate(mgl32.Vec2{5,5}))

	// Make (15, 15) to origin Scale (3x, 3x)
	fmt.Println(canvas.NewTransform().Pivot(mgl32.Vec2{15,15}, func(t *canvas.Transform) {
		t.Scale(mgl32.Vec2{3,3})
	}))
}