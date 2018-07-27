package softcanvas

import (
	"github.com/go-gl/mathgl/mgl32"
	"math"
	"sync"
	"sync/atomic"
	"github.com/iamGreedy/canvas"
	"image"
)

type (
	claabuf struct {
		w, h   int
		prc    float32
		iprc   int32
		rect image.Rectangle
		buffer []cell
	}
	cell struct {
		cover int32
		area  int32
	}
)

func newCLAA(prc Precision, width, height int) *claabuf {
	return &claabuf{
		w:      width,
		h:      height,
		prc:    float32(prc),
		iprc:   int32(prc),
		buffer: make([]cell, width*height),

	}
}

func (s *claabuf) data(q *canvas.Path) {
	s.rect = s.rect.Union(q.Rect)
	//
	wg := new(sync.WaitGroup)
	wg.Add(len(q.Data) - 1)
	for i := 1; i < len(q.Data); i++ {
		p0, p1 := q.Data[i-1], q.Data[i]
		if math.IsNaN(float64(p0[0])) || math.IsNaN(float64(p1[0])) {
			wg.Done()
			continue
		}
		// Points,
		np0 := normalize3(s.prc, p0)
		np1 := normalize3(s.prc, p1)
		// Points, normalize > split > lining
		go func() {
			s.splitsv(np0, np1)
			wg.Done()
		}()
	}
	wg.Wait()
}
func (s *claabuf) splitsv(np0, np1 mgl32.Vec2) {
	//if math.Abs(float64(np0[1] - np1[1])) < float64(1 / s.prc){
	//	return
	//}
	var res = make([]mgl32.Vec2, 0, int32(math.Abs(float64(np1[0]-np0[0]))))
	var unit = 1 / float32(s.prc)
	var tmpx, tmpy, delta = np0[0], np0[1], (np1[1] - np0[1]) / (np1[0] - np0[0]) * unit
	if np0[0] > np1[0] {
		for ; tmpx >= np1[0]; tmpx, tmpy = tmpx-unit, tmpy-delta {
			res = append(res, normalize2(s.prc, mgl32.Vec2{tmpx, tmpy}))
		}
	} else {
		for ; tmpx <= np1[0]; tmpx, tmpy = tmpx+unit, tmpy+delta {
			res = append(res, normalize2(s.prc, mgl32.Vec2{tmpx, tmpy}))
		}
	}
	res = append(res, np1)
	for i := 1; i < len(res); i++ {
		s.splitsh(res[i-1], res[i])
	}
}
func (s *claabuf) splitsh(np0, np1 mgl32.Vec2) {

	//var res = list.New()
	var res = make([]mgl32.Vec2, 0, int32(math.Abs(float64(np1[1]-np0[1]))))
	var unit = 1 / float32(s.prc)
	var tmpx, tmpy, delta = np0[0], np0[1], (np1[0] - np0[0]) / (np1[1] - np0[1]) * unit
	if np0[1] > np1[1] {
		for ; tmpy >= np1[1]; tmpx, tmpy = tmpx-delta, tmpy-unit {
			//res.PushBack(normalize2(s.prc, mgl32.Vec2{tmpx, tmpy}))
			res = append(res, normalize2(s.prc, mgl32.Vec2{tmpx, tmpy}))
		}
	} else {
		for ; tmpy <= np1[1]; tmpx, tmpy = tmpx+delta, tmpy+unit {
			//res.PushBack(normalize2(s.prc, mgl32.Vec2{tmpx, tmpy}))
			res = append(res, normalize2(s.prc, mgl32.Vec2{tmpx, tmpy}))
		}
	}
	//res.PushBack(np1)
	res = append(res, np1)
	//for p, c := res.Front(), res.Front().Next(); c != nil; p, c = c, c.Next(){
	//	s.line(p.Value.(mgl32.Vec2), c.Value.(mgl32.Vec2),)
	//}

	for i := 1; i < len(res); i++ {
		s.line(res[i-1], res[i])
	}
}
func (s *claabuf) line(p0, p1 mgl32.Vec2) {
	cx, cy := int((p0[0]+p1[0])/2), int((p0[1]+p1[1])/2)
	if cy >= s.h || cx >= s.w {
		return
	}
	cover := int32((p1[1] - p0[1]) * float32(s.prc))
	len0 := int32((p0[0] - float32(cx)) * float32(s.prc))
	len1 := int32((p1[0] - float32(cx)) * float32(s.prc))
	area := int32((len0 + len1) * cover)
	//
	offset := s.w*cy + cx
	atomic.AddInt32(&s.buffer[offset].area, area)
	atomic.AddInt32(&s.buffer[offset].cover, cover)
}
func (s *claabuf) effective(c, a int32) uint32 {
	return s.toRange(s.iprc*c - a/2)
}
func (s *claabuf) toRange(value int32) uint32 {
	m := 16 / s.iprc
	res := value * m * m
	if res < 0 {
		res = -res
	}
	if res > math.MaxUint8 {
		res = math.MaxUint8
	}
	return uint32(res)
}
func (s *claabuf) clear() {
	for x := s.rect.Min.X; x < s.rect.Max.X; x++ {
		for y := s.rect.Min.Y; y < s.rect.Max.Y; y++ {
			offset := x + y * s.w
			s.buffer[offset] = cell{}
		}
	}
	//for i, c := range s.buffer {
	//	if c.cover != 0 || c.area != 0 {
	//		s.buffer[i].area = 0
	//		s.buffer[i].cover = 0
	//	}
	//}
}

func normalize3(p float32, v mgl32.Vec3) mgl32.Vec2 {
	return mgl32.Vec2{
		float32(math.Round(float64(v[0]*p))) / p,
		float32(math.Round(float64(v[1]*p))) / p,
	}
}
func normalize2(p float32, v mgl32.Vec2) mgl32.Vec2 {
	return mgl32.Vec2{
		float32(math.Round(float64(v[0]*p))) / p,
		float32(math.Round(float64(v[1]*p))) / p,
	}
}
