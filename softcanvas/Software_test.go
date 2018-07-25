package softcanvas

import (
	"github.com/iamGreedy/canvas"
	"testing"
	// test package already have this
	// _ "github.com/iamGreedy/canvas/softcanvas"
	"github.com/iamGreedy/commons/colors"
	"github.com/iamGreedy/psvg"
	"golang.org/x/image/vector"
	"strings"

	"github.com/go-gl/mathgl/mgl32"
	"image"
	"image/draw"
	"github.com/fogleman/gg"
	"image/color"
	"runtime"
)

var testcase *psvg.Renderer

func init() {

	//
	var err error
	testcase, err = psvg.NewRendererFromReader(strings.NewReader(`
			M362.778,57.968
			c-0.287-0.378-0.565-0.759-0.857-1.136
			c-6.016-7.773-12.405-14.588-19.272-20.518
			c-29.797-25.728-68.582-34.79-124.728-33.123
			c-48.682,1.447-100.473,9.906-134.337,39.903
			c-7.63,6.758-14.343,14.617-19.913,23.729
			c-0.726,1.188-1.439,2.391-2.126,3.622
			c-2.921,5.239-5.633,10.771-8.112,16.638
			c-3.476,8.224-6.49,17.108-8.95,26.793
			c-4.767,18.77-7.463,40.533-7.462,66.257
			c0.002,45.133,8.866,67.528,8.332,110.879
			c-0.035,2.846-0.11,5.782-0.231,8.821
			c-0.204,5.119-0.383,10.249-0.543,15.375
			c-1.653,53.107-1.062,105.862-1.499,142.036
			c-0.401,33.204,14.646,62.704,41.845,84.433
			c10.752,8.59,23.402,15.965,37.752,21.871
			c25.113,10.337,55.418,16.186,89.844,16.186
			c50.265,0,87.456-9.652,114.684-24.336
			c11.459-6.178,21.149-13.249,29.308-20.867
			c20.359-19.008,31.17-41.422,36.009-61.896
			c11.47-48.523,9.966-84.08,4.831-158.371
			c-0.103-1.496-0.207-3.002-0.313-4.529
			c-0.588-8.453-1.022-15.947-1.339-22.763
			c-1.733-37.343,0.064-54.317-0.479-96.937
			c-0.463-36.271-3.195-63.161-9.306-85.047
			c-2.776-9.942-6.244-18.858-10.521-27.142
			C371.785,70.854,367.594,64.312,362.778,57.968
			z`))
	if err != nil {
		panic(err)
	}
}

func TestSoftware_Query_Simple(t *testing.T) {
	//
	suf, err := canvas.NewSurface(Driver, 16, 16,
		//canvas.Clear{0, 0, 0, 0},
	)
	if err != nil {
		panic(err)
	}
	if res := suf.Option(canvas.Clear(color.RGBAModel.Convert(colors.HTML.White).(color.RGBA))); canvas.IsFail(res) {
		panic(res)
	}

	suf.Query(canvas.NewPath().Fill(func(i canvas.InnerPath) {
		i.MoveTo(mgl32.Vec2{4,4})
		i.LineTo(mgl32.Vec2{12,12})
		i.LineTo(mgl32.Vec2{12,4})
		i.CloseTo()
	}), canvas.NewColorShader(colors.HTML.Black), nil)
	//
	canvas.Capture("out.png", suf)
}

func TestSoftware_Query(t *testing.T) {
	//
	suf, err := canvas.NewSurface(Driver, 1024, 1024,
	)
	if err != nil {
		panic(err)
	}
	if res := suf.Option(canvas.Clear{0, 0, 0, 0}); canvas.IsFail(res) {
		panic(res)
	}

	suf.Query(canvas.NewPath().Fill(func(i canvas.InnerPath) {
		testcase.Render(i)
	}), canvas.NewColorShader(colors.HTML.Black), nil)
	//
	canvas.Capture("out.png", suf)
}



func BenchmarkCanvas(b *testing.B) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	b.StopTimer()
	suf, err := canvas.NewSurface(Driver, 1024, 1024,
		canvas.Src)
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		if r := suf.Option(canvas.Clear{0, 0, 0, 255}); canvas.IsFail(r) {
			panic(r)
		}
		b.StartTimer()
		suf.Query(canvas.NewPath().Fill(func(i canvas.InnerPath) {
			testcase.Render(i)
		}), canvas.NewColorShader(colors.HTML.White), nil)
		b.StopTimer()
	}
}

type psvglayer struct {
	r *vector.Rasterizer
}

func (s psvglayer) MoveTo(to mgl32.Vec2) {
	s.r.MoveTo(to[0], to[1])
}
func (s psvglayer) LineTo(to mgl32.Vec2) {
	s.r.LineTo(to[0], to[1])
}
func (s psvglayer) QuadTo(p0, to mgl32.Vec2) {
	s.r.QuadTo(p0[0], p0[1], to[0], to[1])
}
func (s psvglayer) CubeTo(p0, p1, to mgl32.Vec2) {
	s.r.CubeTo(p0[0], p0[1], p1[0], p1[1], to[0], to[1])
}
func (s psvglayer) CloseTo() {
	s.r.ClosePath()
}

func BenchmarkVector(b *testing.B) {
	b.StopTimer()
	tmp := vector.NewRasterizer(1024, 1024)
	img := image.NewRGBA(image.Rect(0, 0, 1024, 1024))
	rst := psvglayer{
		r: tmp,
	}
	tmp.DrawOp = draw.Src
	//
	for i := 0; i < b.N; i++ {
		draw.Draw(img, img.Rect, image.NewUniform(colors.HTML.White), image.ZP, draw.Src)
		b.StartTimer()
		testcase.Render(rst)
		rst.r.Draw(img, img.Rect, image.NewUniform(colors.HTML.Black), image.ZP)
		b.StopTimer()
	}

}

type gglayer struct {
	ctx *gg.Context
}
func (s gglayer) MoveTo(to mgl32.Vec2) {
	s.ctx.MoveTo(float64(to[0]), float64(to[1]))
}
func (s gglayer) LineTo(to mgl32.Vec2) {
	s.ctx.LineTo(float64(to[0]), float64(to[1]))
}
func (s gglayer) QuadTo(p0, to mgl32.Vec2) {
	s.ctx.QuadraticTo(float64(p0[0]), float64(p0[1]), float64(to[0]), float64(to[1]))
}
func (s gglayer) CubeTo(p0, p1, to mgl32.Vec2) {
	s.ctx.CubicTo(float64(p0[0]), float64(p0[1]), float64(p1[0]), float64(p1[1]), float64(to[0]), float64(to[1]))
}
func (s gglayer) CloseTo() {
	s.ctx.ClosePath()
}

func BenchmarkGG(b *testing.B) {
	b.StopTimer()
	ctx := gg.NewContext(1024, 1024)
	ctx.SetColor(colors.HTML.Black)
	lgg := gglayer{
		ctx:ctx,
	}
	for i := 0; i < b.N; i++ {
		ctx.Clear()
		b.StartTimer()
		testcase.Render(lgg)
		ctx.Fill()
		b.StopTimer()
	}
}

