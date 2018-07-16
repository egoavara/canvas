package canvas

import (
	"testing"
	"sync"
	"fmt"
)

func BenchmarkToSurfaceType(b *testing.B) {

	b.StopTimer()
	new(sync.Once).Do(func() {
		fmt.Println("BenchmarkToSurfaceType")
	})
	cases := []string{"Legacy", "OpenGL", "invalid case"}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ToSurfaceType(cases[i % 3])
	}
}
func BenchmarkSurfaceType_String(b *testing.B) {

	b.StopTimer()
	new(sync.Once).Do(func() {
		fmt.Println("BenchmarkSurfaceType_String")
	})

	cases := []SurfaceType{SurfaceTypeOpenGL, SurfaceTypeSoftware, SurfaceTypeInvalid}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		cases[i % 3].String()
	}
}