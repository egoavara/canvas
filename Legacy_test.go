package canvas

import (
	"testing"
	"image/color"
	"github.com/go-gl/mathgl/mgl32"
	"golang.org/x/image/vector"
	"image"
	"runtime"
	"github.com/golang/freetype/raster"
	"golang.org/x/image/math/fixed"
	"image/draw"
)

func BenchmarkSoftware_Query(b *testing.B) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//
	surf, err := NewSurface(1024, 1024, SurfaceTypeSoftware)
	if err != nil {
		panic(err)
	}
	for i := 0; i < b.N; i++ {
		surf.Flush()
		b.StartTimer()
		surf.Query(NewPath().Fill(func(i InnerPath) {
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
		}), NewColorShader(color.Black), NewTransform())
		b.StopTimer()
	}
}
func BenchmarkVector(b *testing.B) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//
	v := vector.NewRasterizer(1024, 1024)
	res := image.NewRGBA(image.Rect(0,0,1024,1024))
	const padding = 4
	const radius = 28
	//
	for i := 0; i < b.N; i++ {

		b.StartTimer()
		v.MoveTo(
			padding + radius,
			padding)
		v.QuadTo(
			
				padding,
				padding,
			
				padding,
				padding + radius)
		v.LineTo(
			padding,
			1024 - padding - radius)
		v.QuadTo(
			
				padding,
				1024 - padding,
			
				padding + radius,
				1024 - padding)
		v.LineTo(
			1024 - padding - radius,
			1024 - padding)
		v.QuadTo(
			
				1024 - padding,
				1024 - padding,
			
				1024 - padding,
				1024 - padding - radius)
		v.LineTo(
			1024 - padding,
			padding + radius)
		v.QuadTo(
			1024 - padding,
			padding, 
			1024 - padding - radius,
			padding)
		v.ClosePath()
		v.Draw(res, res.Rect, image.NewUniform(color.Black), image.ZP)
		b.StopTimer()
	}
}
func BenchmarkFreetype(b *testing.B) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//

	v := raster.NewRasterizer(1024, 1024)
	black := image.NewRGBA(image.Rect(0,0,1024,1024))
	draw.Draw(black, black.Rect, image.NewUniform(color.Black), image.ZP, draw.Src)
	painter := raster.NewRGBAPainter(black)
	const padding = 4
	const radius = 28
	//
	for i := 0; i < b.N; i++ {

		b.StartTimer()
		v.Start(fixed.P(
			padding + radius,
			padding))
		v.Add2(
			fixed.P(
				padding,
				padding),
			fixed.P(
				padding,
				padding + radius))
		v.Add1(fixed.P(
			padding,
			1024 - padding - radius))
		v.Add2(
			fixed.P(
				padding,
				1024 - padding),
			fixed.P(
				padding + radius,
				1024 - padding))
		v.Add1(fixed.P(
			1024 - padding - radius,
			1024 - padding))
		v.Add2(
			fixed.P(
				1024 - padding,
				1024 - padding),
			fixed.P(
				1024 - padding,
				1024 - padding - radius))
		v.Add1(fixed.P(
			1024 - padding,
			padding + radius))
		v.Add2(fixed.P(
			1024 - padding,
			padding), fixed.P(
			1024 - padding - radius,
			padding))
		v.Rasterize(painter)
		b.StopTimer()
	}
}