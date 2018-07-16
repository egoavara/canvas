package main

import (
	"golang.org/x/image/math/fixed"
	"runtime"
	"github.com/golang/freetype/raster"
	"image"
	"os"
	"image/png"
	"image/color"
	"golang.org/x/image/draw"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//
	out := image.NewRGBA(image.Rect(0,0,1024, 1024))
	v := raster.NewRasterizer(1024, 1024)
	//v.
	//black := image.NewRGBA(image.Rect(0,0,1024,1024))
	//draw.Draw(black, black.Rect, image.NewUniform(color.Black), image.ZP, draw.Src)
	//painter := raster.NewRGBAPainter(black)
	const padding = 4
	const radius = 28
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


	painter := raster.NewRGBAPainter(out)
	painter.SetColor(color.Black)
	painter.Op = draw.Src

	v.Rasterize(painter)
	f, err := os.Create("out.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, out)
}
