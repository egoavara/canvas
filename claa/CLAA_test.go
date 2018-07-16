package claa

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/iamGreedy/canvas"
	"image"
	"image/png"
	"os"
	"testing"
	"runtime"
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
	claa := NewCLAA(X4, width, height)
	data := canvas.NewPath().Fill(func(i canvas.InnerPath) {
		const padding = 16
		i.MoveTo(mgl32.Vec2{padding, padding})
		i.LineTo(mgl32.Vec2{padding, height - padding})
		i.QuadTo(mgl32.Vec2{width - padding, height - padding}, mgl32.Vec2{width - padding, padding})
		i.CloseTo()
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
