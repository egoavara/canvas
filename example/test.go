package main

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/iamGreedy/canvas"
	"github.com/iamGreedy/canvas/claa"
)

func main() {
	//surf, err := canvas.NewSurface(1024, 1024, canvas.SurfaceTypeSoftware)
	//if err != nil {
	//	panic(err)
	//}
	//keys := colors.HTML.Keys()
	//for k, p := range Paths {
	//	surf.Query(p, canvas.NewColorShader(colors.HTML.Find(keys[int(rand.Uint32())%len(keys)])), nil)
	//	canvas.Capture(k+".png", surf)
	//	surf.Clear()
	//}
	//small, err := canvas.NewSurface(32, 32, canvas.SurfaceTypeSoftware)
	//if err != nil {
	//	panic(err)
	//}
	//small.Query(canvas.NewPath().Fill(func(i canvas.InnerPath) {
	//	i.MoveTo(mgl32.Vec2{2,2})
	//	i.LineTo(mgl32.Vec2{30,30})
	//	i.LineTo(mgl32.Vec2{30,2})
	//	i.CloseTo()
	//
	//	i.MoveTo(mgl32.Vec2{2,2})
	//	i.LineTo(mgl32.Vec2{30,2})
	//	i.LineTo(mgl32.Vec2{2,30})
	//	i.CloseTo()
	//}), canvas.NewColorShader(colors.HTML.Black), nil)
	//canvas.Capture("out.png", small)
	//var a canvas.SortedSignIgnoreFixedList
	//a.Append(fixed.Int26_6(192), )
	//a.Append(fixed.Int26_6(-128), )
	//fmt.Println(a)
	//raster.NewRasterizer()
	c := []mgl32.Vec3{
		{0.5, 0.6, 1.},
		{3.6, 2.8, 1.},
		{2.1, 0.1, 1.},
		{0.5, 0.6, 1.},
	}
	//claa.Normalize(claa.X4, c...)
	//v0 := claa.Split(c[0], c[1])
	//claa.Normalize(claa.X4, v0...)
	//fmt.Println(c)
	//fmt.Println(v0)
	pre := claa.X4
	res := claa.Cellize(3, pre, c)
	//res := claa.AlignGrid(claa.X4, c)
	for i, p := range res {
		fmt.Print(i, " : ")
		for _, c := range p {
			fmt.Printf("%-40v ", c)
		}
		fmt.Println()
	}
	var Pix = make([]uint8, 12)
	claa.Raster(pre, res, Pix, 4)
	for y := 2; y >= 0; y -- {
		for x := 0; x < 4; x ++ {
			fmt.Printf("%3d ", Pix[y * 4 + x])
		}
		fmt.Println()
	}
	fmt.Println(Pix)
}

var Paths = map[string]*canvas.Path{
	"CW": canvas.NewPath().Fill(func(i canvas.InnerPath) {
		const padding = 4
		i.MoveTo(mgl32.Vec2{0 + padding, 0 + padding})
		i.LineTo(mgl32.Vec2{0 + padding, 1024 - padding})
		i.LineTo(mgl32.Vec2{1024 - padding, 1024 - padding})
		i.CloseTo()
	}),
	"CCW": canvas.NewPath().Fill(func(i canvas.InnerPath) {
		const padding = 4
		i.MoveTo(mgl32.Vec2{0 + padding, 0 + padding})
		i.LineTo(mgl32.Vec2{1024 - padding, 0 + padding})
		i.LineTo(mgl32.Vec2{1024 - padding, 1024 - padding})
		i.CloseTo()
	}),

	"Intersect": canvas.NewPath().Fill(func(i canvas.InnerPath) {
		const padding = 4

		i.MoveTo(mgl32.Vec2{0 + padding, 0 + padding})
		i.LineTo(mgl32.Vec2{1024 - padding, 0 + padding})
		i.LineTo(mgl32.Vec2{1024 - padding, 1024 - padding})
		i.CloseTo()
		//
		i.MoveTo(mgl32.Vec2{0 + padding, 0 + padding})
		i.LineTo(mgl32.Vec2{0 + padding, 1024 - padding})
		i.LineTo(mgl32.Vec2{1024 - padding, 0 + padding})
		i.CloseTo()
	}),
	"Arc": canvas.NewPath().Fill(func(i canvas.InnerPath) {
		const padding = 4
		i.MoveTo(mgl32.Vec2{0 + padding, 1024 - padding})
		i.LineTo(mgl32.Vec2{1024 - padding, 1024 - padding})
		i.QuadTo(mgl32.Vec2{1024 - padding, 0 + padding}, mgl32.Vec2{0 + padding, 0 + padding})
		i.CloseTo()
	}),
	"SmoothCube": canvas.NewPath().Fill(func(i canvas.InnerPath) {
		const padding = 4
		const radius = 28
		i.MoveTo(mgl32.Vec2{
			padding + radius,
			padding})
		i.QuadTo(
			mgl32.Vec2{
				padding,
				padding},
			mgl32.Vec2{
				padding,
				padding + radius})
		i.LineTo(mgl32.Vec2{
			padding,
			1024 - padding - radius})
		i.QuadTo(
			mgl32.Vec2{
				padding,
				1024 - padding},
			mgl32.Vec2{
				padding + radius,
				1024 - padding})
		i.LineTo(mgl32.Vec2{
			1024 - padding - radius,
			1024 - padding})
		i.QuadTo(
			mgl32.Vec2{
				1024 - padding,
				1024 - padding},
			mgl32.Vec2{
				1024 - padding,
				1024 - padding - radius})
		i.LineTo(mgl32.Vec2{
			1024 - padding,
			padding + radius})
		i.QuadTo(mgl32.Vec2{
			1024 - padding,
			padding}, mgl32.Vec2{
			1024 - padding - radius,
			padding})
		i.CloseTo()
	}),
}
