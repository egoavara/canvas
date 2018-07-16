package claa

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"math"
	"sort"
	"sync"
)

type cell struct {
	L0, L1 mgl32.Vec3
	X      int32
	Cover  int32
	Area   int32
}

func (s cell) String() string {
	//return fmt.Sprintf("cell(x : %4d, %v - %v)", s.X, s.L0.Vec2(), s.L1.Vec2())
	return fmt.Sprintf("cell(x : %4d, cover : %4d, area : %4d)", s.X, s.Cover, s.Area)
}
func Raster(p Precision, cell []cellstream, dst []uint8, stride int32) {
	width := int32(p)
	for y, stream := range cell {
		var x, cover, area int32 = 0, 0, 0
		var prev int32 = 0

		for _, v := range stream {
			if x != v.X {
				prev = effective(width, cover+prev, area)
				dst[stride*int32(y)+x] = toRange(prev, width)
				x = v.X
				cover = 0
				area = 0
			}
			cover += v.Cover
			area += v.Area

		}
		prev = effective(width, cover+prev, area)
		dst[stride*int32(y)+x] = toRange(prev, width)
	}
}
func Cellize(height uint32, p Precision, points []mgl32.Vec3) (res []cellstream) {
	points = AlignGrid(p, points)
	res = make([]cellstream, height)
	for i := 1; i < len(points); i++ {
		l0, l1 := points[i-1], points[i]
		cx, cy := int32((l0[0]+l1[0])/2), int32((l0[1]+l1[1])/2)
		cover := int32((l1[1] - l0[1]) * float32(p))
		len0 := int32((l0[0] - float32(cx)) * float32(p))
		len1 := int32((l1[0] - float32(cx)) * float32(p))
		area := int32((len0 + len1) * cover)

		//
		//
		res[cy] = append(res[cy], &cell{
			L0:    l0,
			L1:    l1,
			X:     cx,
			Cover: cover,
			Area:  area,
		})
	}
	for _, l := range res {
		sort.Sort(l)
	}
	return
}
func AlignGrid(p Precision, points []mgl32.Vec3) (res []mgl32.Vec3) {
	from := 0
	wg := new(sync.WaitGroup)
	ch := make(chan []mgl32.Vec3)
	mtx := new(sync.Mutex)
	go func() {
		mtx.Lock()
		defer mtx.Unlock()
		// result list build goroutine
		for l := range ch{
			res = append(res, l...)
		}

	}()
	for to := 0; to < len(points); to++ {
		if math.IsNaN(float64(points[to][0])) {
			wg.Add(1)
			if to-from > 3 {
				go func(from, to int) {
					// pallarel works
					ch <- append(alignGrid(p, points[from:to]), mgl32.Vec3{float32(math.NaN()), float32(math.NaN()), float32(math.NaN())})
					wg.Done()
				}(from, to)
			}
			to++
			from = to
			continue
		}
	}
	wg.Wait()
	close(ch)
	mtx.Lock()
	return res
}
func alignGrid(p Precision, points []mgl32.Vec3) (res []mgl32.Vec3) {
	Normalize(p, points...)
	//
	var verticals []mgl32.Vec3
	for i := 0; i < len(points)-1; i++ {
		verticals = append(verticals, points[i])
		verticals = append(verticals, Split(points[i], points[i+1], 0)...)
	}
	verticals = append(verticals, points[len(points)-1])
	Normalize(p, verticals...)
	for i := 0; i < len(verticals)-1; i++ {
		res = append(res, verticals[i])
		res = append(res, Split(verticals[i], verticals[i+1], 1)...)
	}
	res = append(res, verticals[len(verticals)-1])
	Normalize(p, res...)
	return res
}

func Normalize(p Precision, points ...mgl32.Vec3) {
	fp := float32(p)
	for i, p := range points {
		points[i] = mgl32.Vec3{
			float32(math.Floor(float64(p[0]*fp))) / fp,
			float32(math.Floor(float64(p[1]*fp))) / fp,
			1,
		}
	}
	return
}
func SplitSize(l0, l1 float32) int {
	if l0 > l1 {
		l0, l1 = l1, l0
	}
	return int(math.Ceil(float64(l1)) - math.Ceil(float64(l0)))
}

// 0 = vertical
// 1 = horizontal
// l0, l1 is normalized points
func Split(l0, l1 mgl32.Vec3, using int) (res []mgl32.Vec3) {
	notusing := using ^ 1
	res = make([]mgl32.Vec3, SplitSize(l0[notusing], l1[notusing]))
	if len(res) == 0 {
		return
	}
	var idx, didx int
	if l0[notusing] > l1[notusing] {
		l0, l1 = l1, l0
		idx, didx = len(res)-1, -1
	} else {
		idx, didx = 0, 1
	}
	d := (l1[using] - l0[using]) / (l1[notusing] - l0[notusing])
	a := l0[using] + d*float32(math.Ceil(float64(l0[notusing]))-float64(l0[notusing]))
	for b := float32(math.Ceil(float64(l0[notusing]))); b < float32(math.Ceil(float64(l1[notusing]))); b += 1 {
		var temp mgl32.Vec3
		temp[using] = a
		temp[notusing] = b
		temp[2] = 1
		res[idx] = temp
		idx += didx
		a += d
	}
	if didx == -1 {
		if res[0] == l1 {
			res = res[1:]
		}
		if len(res) > 0 && res[len(res)-1] == l0 {
			res = res[:len(res)-1]
		}
	} else {
		if res[0] == l0 {
			res = res[1:]
		}
		if len(res) > 0 && res[len(res)-1] == l1 {
			res = res[:len(res)-1]
		}
	}

	return
}

type cellstream []*cell

func (s cellstream) Len() int {
	return len(s)
}
func (s cellstream) Less(i, j int) bool {
	return s[i].X < s[j].X
}
func (s cellstream) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
