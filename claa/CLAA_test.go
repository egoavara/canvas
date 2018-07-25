package claa

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/iamGreedy/canvas"
	"image"
	"image/png"
	"os"
	"testing"
	"runtime"
	"github.com/iamGreedy/psvg"
	"strings"
)

func TestCLAA(t *testing.T) {

	const width, height = 1024, 1024
	// ready
	f, err := os.Create("test.png")
	if err != nil {
		t.Fatal(err)
		return
	}
	defer f.Close()
	aimg := image.NewAlpha(image.Rect(0, 0, width, height))
	claa := NewCLAA(X16, width, height)
	psrd := MakePath()
	data := canvas.NewPath().Fill(func(i canvas.InnerPath) {
		psrd.Render(i)
	}).Data
	// start
	claa.Data(RAW, data...)
	claa.Result(aimg.Pix, aimg.Stride, 1, 0)
	// save result
	png.Encode(f, aimg)
}

func BenchmarkAlignGrid(b *testing.B) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	b.StopTimer()
	const width, height = 1024, 1024
	data := canvas.NewPath().Fill(func(i canvas.InnerPath) {
		const padding = 16
		i.MoveTo(mgl32.Vec2{padding, padding})
		i.LineTo(mgl32.Vec2{padding, height - padding})
		//i.LineTo(mgl32.Vec2{width - padding, height - padding})
		//i.LineTo(mgl32.Vec2{width - padding, padding})
		i.QuadTo(mgl32.Vec2{width - padding, height - padding}, mgl32.Vec2{width - padding, padding})
		i.CloseTo()
	}).Data
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		AlignGrid(X4, data)
		b.StopTimer()
	}
}
func BenchmarkCLAARAW(b *testing.B) {
	b.StopTimer()
	const width, height = 1024, 1024

	claa := NewCLAA(X4, width, height)
	// ready
	aimg := image.NewAlpha(image.Rect(0, 0, width, height))
	for i := 0; i < b.N; i++ {
		claa.Clear()
		b.StartTimer()
		data := canvas.NewPath().Fill(func(i canvas.InnerPath) {
			const padding = 16
			i.MoveTo(mgl32.Vec2{padding, padding})
			i.LineTo(mgl32.Vec2{padding, height - padding})
			//i.LineTo(mgl32.Vec2{width - padding, height - padding})
			//i.LineTo(mgl32.Vec2{width - padding, padding})
			i.QuadTo(mgl32.Vec2{width - padding, height - padding}, mgl32.Vec2{width - padding, padding})
			i.CloseTo()
		}).Data
		// start
		claa.Data(RAW, data...)
		claa.Result(aimg.Pix, aimg.Stride, 1, 0)
		b.StopTimer()
	}
}
func BenchmarkCLAAALIGNED(b *testing.B) {
	b.StopTimer()
	const width, height = 1024, 1024

	claa := NewCLAA(X4, width, height)
	data := canvas.NewPath().Fill(func(i canvas.InnerPath) {
		const padding = 16
		i.MoveTo(mgl32.Vec2{padding, padding})
		i.LineTo(mgl32.Vec2{padding, height - padding})
		//i.LineTo(mgl32.Vec2{width - padding, height - padding})
		//i.LineTo(mgl32.Vec2{width - padding, padding})
		i.QuadTo(mgl32.Vec2{width - padding, height - padding}, mgl32.Vec2{width - padding, padding})
		i.CloseTo()
	}).Data
	data = AlignGrid(X4, data)
	// ready
	aimg := image.NewAlpha(image.Rect(0, 0, width, height))
	for i := 0; i < b.N; i++ {
		claa.Clear()
		b.StartTimer()
		// start
		claa.Data(ALIGNED, data...)
		claa.Result(aimg.Pix, aimg.Stride, 1, 0)
		b.StopTimer()
	}
}
func BenchmarkCLAASyncGoRoutine(b *testing.B) {
	const width, height = 1024, 1024
	claa := NewCLAA(X16, width, height)
	data := canvas.NewPath().Fill(func(i canvas.InnerPath) {
		const padding = 16
		i.MoveTo(mgl32.Vec2{padding, padding})
		i.LineTo(mgl32.Vec2{padding, height - padding})
		//i.LineTo(mgl32.Vec2{width - padding, height - padding})
		//i.LineTo(mgl32.Vec2{width - padding, padding})
		i.QuadTo(mgl32.Vec2{width - padding, height - padding}, mgl32.Vec2{width - padding, padding})
		i.CloseTo()
	}).Data
	data = AlignGrid(X4, data)
	// ready
	aimg := image.NewAlpha(image.Rect(0, 0, width, height))
	b.Run("Sync", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			claa.Clear()
			b.StartTimer()
			// start
			claa.Data(ALIGNED, data...)
			claa.SyncResult(aimg.Pix, aimg.Stride, 1, 0)
			b.StopTimer()
		}
	})
	b.Run("Goroutine", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			claa.Clear()
			b.StartTimer()
			// start
			claa.Data(ALIGNED, data...)
			claa.Result(aimg.Pix, aimg.Stride, 1, 0)
			b.StopTimer()
		}
	})
}

func MakePath() *psvg.Renderer {
	p, err := psvg.NewRendererFromReader(strings.NewReader(`
M362.778,57.968
c-0.287-0.378-0.565-0.759-0.857-1.136c-6.016-7.773-12.405-14.588-19.272-20.518
	c-29.797-25.728-68.582-34.79-124.728-33.123c-48.682,1.447-100.473,9.906-134.337,39.903c-7.63,6.758-14.343,14.617-19.913,23.729
	c-0.726,1.188-1.439,2.391-2.126,3.622c-2.921,5.239-5.633,10.771-8.112,16.638c-3.476,8.224-6.49,17.108-8.95,26.793
	c-4.767,18.77-7.463,40.533-7.462,66.257c0.002,45.133,8.866,67.528,8.332,110.879c-0.035,2.846-0.11,5.782-0.231,8.821
	c-0.204,5.119-0.383,10.249-0.543,15.375c-1.653,53.107-1.062,105.862-1.499,142.036c-0.401,33.204,14.646,62.704,41.845,84.433
	c10.752,8.59,23.402,15.965,37.752,21.871c25.113,10.337,55.418,16.186,89.844,16.186c50.265,0,87.456-9.652,114.684-24.336
	c11.459-6.178,21.149-13.249,29.308-20.867c20.359-19.008,31.17-41.422,36.009-61.896c11.47-48.523,9.966-84.08,4.831-158.371
	c-0.103-1.496-0.207-3.002-0.313-4.529c-0.588-8.453-1.022-15.947-1.339-22.763c-1.733-37.343,0.064-54.317-0.479-96.937
	c-0.463-36.271-3.195-63.161-9.306-85.047c-2.776-9.942-6.244-18.858-10.521-27.142C371.785,70.854,367.594,64.312,362.778,57.968z`))
	if err != nil {
		panic(err)
	}
	return p
}