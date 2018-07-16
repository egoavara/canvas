package canvas

import (
	"github.com/go-gl/mathgl/mgl32"
	"golang.org/x/image/math/fixed"
	"github.com/iamGreedy/canvas/fx32"
)

type sortedSignIgnoreF32List []float32

func (s *sortedSignIgnoreF32List) Append(f32s ...float32) {
	for _, f := range f32s {
		s.append(f)
	}
}
func (s *sortedSignIgnoreF32List) append(f32 float32) {
	if len(*s) == 0{
		*s = sortedSignIgnoreF32List{f32}
		return
	}
	if len(*s) == 1{
		if mgl32.Abs((*s)[0]) > mgl32.Abs(f32){
			*s = sortedSignIgnoreF32List{f32, (*s)[0]}
		}else {
			*s = sortedSignIgnoreF32List{(*s)[0], f32}
		}
		return
	}
	toi := 0
	for i, f := range *s {
		if mgl32.Abs(f) > mgl32.Abs(f32){
			break
		}
		toi = i
	}
	toi += 1

	*s = append(*s, 0)
	for i := len(*s) - 1; i >= toi + 1; i-- {
		(*s)[i] = (*s)[i-1]
	}
	(*s)[toi] = f32
}

type SortedSignIgnoreFixedList []fixed.Int26_6

func (s *SortedSignIgnoreFixedList) Append(f32s ...fixed.Int26_6) {
	for _, f := range f32s {
		s.append(f)
	}
}
func (s *SortedSignIgnoreFixedList) append(f32 fixed.Int26_6) {
	if len(*s) == 0{
		*s = SortedSignIgnoreFixedList{f32}
		return
	}
	if len(*s) == 1{
		if fx32.Abs((*s)[0]) > fx32.Abs(f32){
			*s = SortedSignIgnoreFixedList{f32, (*s)[0]}
		}else {
			*s = SortedSignIgnoreFixedList{(*s)[0], f32}
		}
		return
	}
	toi := 0
	for i, f := range *s {
		if fx32.Abs(f) > fx32.Abs(f32){
			break
		}
		toi = i
	}
	toi += 1

	*s = append(*s, 0)
	for i := len(*s) - 1; i >= toi + 1; i-- {
		(*s)[i] = (*s)[i-1]
	}
	(*s)[toi] = f32
}