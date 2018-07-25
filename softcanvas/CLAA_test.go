package softcanvas

import (
	"testing"
	"github.com/go-gl/mathgl/mgl32"
	"fmt"
)

func TestClaa(t *testing.T) {
	const width = 8
	const height = 8
	claa := newCLAA(X2, width, height)
	claa.data(
		mgl32.Vec3{2, 2, 1},
		mgl32.Vec3{6, 6, 1},
		mgl32.Vec3{6, 2, 1},
		mgl32.Vec3{2, 2, 1},
	)
	for y := 0; y < height; y++{
		for x := 0; x < width; x++ {
			offset := y * width + x
			fmt.Printf("(%4d, %4d) ", claa.buffer[offset].cover, claa.buffer[offset].area)
		}
		fmt.Println()
	}
}

func BenchmarkCLAA(b *testing.B) {
	//const width = 1024
	//const height = 1024
	//claa := newCLAA(X2, width, height)
	//for i := 0; i < b.N; i++{
	//
	//}
}