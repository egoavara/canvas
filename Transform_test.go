package canvas

import (
	"testing"
	"github.com/go-gl/mathgl/mgl32"
	"sync"
	"fmt"
	"github.com/iamGreedy/axis"
)

func BenchmarkNewTransfrom(b *testing.B) {
	b.StopTimer()
	new(sync.Once).Do(func() {
		fmt.Println("BenchmarkNewTransfrom")
	})
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		NewTransform()
	}
}
func BenchmarkTransfrom_Rotate(b *testing.B) {
	b.StopTimer()
	new(sync.Once).Do(func() {
		fmt.Println("BenchmarkTransfrom_Rotate")
	})
	var m = NewTransform()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		m.Rotate(45)
	}
}
func BenchmarkTransfrom_Translate(b *testing.B) {
	b.StopTimer()
	var m = NewTransform()

	new(sync.Once).Do(func() {
		fmt.Println("BenchmarkTransfrom_Translate")
	})
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		m.Translate(mgl32.Vec2{1,1})
	}
}
func BenchmarkTransfrom_Scale(b *testing.B) {
	b.StopTimer()
	var m = NewTransform()

	new(sync.Once).Do(func() {
		fmt.Println("BenchmarkTransfrom_Scale")
	})
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		m.Scale(mgl32.Vec2{2,2})
	}
}
func BenchmarkTransfrom_Reflection(b *testing.B) {
	b.StopTimer()
	var m = NewTransform()
	var cases = []axis.Axis{axis.Both, axis.Vertical, axis.Horizontal}
	new(sync.Once).Do(func() {
		fmt.Println("BenchmarkTransfrom_Reflection")
	})
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		m.Reflection(cases[i % 3])
	}
}
func BenchmarkTransfrom_Shear(b *testing.B) {

	b.StopTimer()
	var m = NewTransform()
	var cases = []axis.ExculsiveAxis{axis.ExculsiveVertical, axis.ExculsiveHorizontal}
	new(sync.Once).Do(func() {
		fmt.Println("BenchmarkTransfrom_Shear")
	})
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		m.Shear(3, cases[i % 2])
	}
}
func BenchmarkTransfrom_Pivot(b *testing.B) {
	b.StopTimer()
	var m = NewTransform()

	new(sync.Once).Do(func() {
		fmt.Println("BenchmarkTransfrom_Pivot")
	})
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		m.Pivot(mgl32.Vec2{1,3}, func(t *Transform) {

		})
	}
}
func BenchmarkTransfrom_Mul(b *testing.B) {
	b.StopTimer()
	var a = mgl32.Vec2{1,1}
	var m = NewTransform()

	new(sync.Once).Do(func() {
		fmt.Println("BenchmarkTransfrom_Mul")
	})
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		m.Mul(a)
	}
}
func BenchmarkTransfrom_String(b *testing.B) {
	b.StopTimer()
	var m = NewTransform()

	new(sync.Once).Do(func() {
		fmt.Println("BenchmarkTransfrom_String")
	})
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		m.String()
	}
}