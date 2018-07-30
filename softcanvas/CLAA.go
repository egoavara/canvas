package softcanvas

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/iamGreedy/canvas"
	"image"
	"math"
	"sync"
	"sync/atomic"
)

const blockside = 64
const blocksize = blockside * blockside

type (
	claaManager struct {
		ref *Software
		//
		wait []*claablock
		lck  sync.Mutex
	}
	claablock [blocksize]cell
	//
	claa struct {
		w      int
		prc    float32
		iprc   int32
		man    *claaManager
		rect   image.Rectangle
		buffer []*claablock
	}
	cell struct {
		cover int32
		area  int32
	}
)

func newClaaManager(ref *Software) *claaManager {
	return &claaManager{
		ref: ref,
	}
}
func (s *claaManager) get(rectangle image.Rectangle) *claa {
	var needw, needh, left int
	dx, dy := rectangle.Dx(), rectangle.Dy()
	if needw, left = dx/blockside, dx%blockside; left != 0 {
		needw += 1
	}
	if needh, left = dy/blockside, dy%blockside; left != 0 {
		needh += 1
	}
	//
	totalneed := needw * needh
	s.lck.Lock()
	defer s.lck.Unlock()
	if remain := len(s.wait) - totalneed; remain < 0 {
		temp := make([]*claablock, -remain)
		for i := range temp {
			temp[i] = new(claablock)
		}
		temp = append(temp, s.wait...)
		s.wait = nil
		return &claa{
			w:      needw,
			rect:   rectangle,
			buffer: temp,
			man:    s,
			prc:    float32(s.ref.precision),
			iprc:   int32(s.ref.precision),
		}
	}
	temp := s.wait[:totalneed]
	s.wait = s.wait[totalneed:]
	return &claa{
		w:      needw,
		rect:   rectangle,
		buffer: temp,
		man:    s,
		prc:    float32(s.ref.precision),
		iprc:   int32(s.ref.precision),
	}
}
func (s *claaManager) put(c *claa) {
	for _, b := range c.buffer {
		for _, i := range b {
			i.area = 0
			i.cover = 0
		}
	}
	s.lck.Lock()
	defer s.lck.Unlock()
	s.wait = append(s.wait, c.buffer...)
	if mx := s.max(); len(s.wait) > mx {
		s.wait = s.wait[:mx]
	}
}
func (s *claaManager) max() int {
	return ((s.ref.size.X/blockside + 1) * (s.ref.size.Y/blockside + 1)) * 2
}

//
func (s *claa) offset(x, y int) (b, i int) {
	if x < s.rect.Min.X || s.rect.Max.X <= x {
		return -1, -1
	}
	if y < s.rect.Min.Y || s.rect.Max.Y <= y {
		return -1, -1
	}
	x -= s.rect.Min.X
	y -= s.rect.Min.Y
	//
	bx := x / blockside
	by := y / blockside
	b = bx + by*s.w
	ix := x % blockside
	iy := y % blockside
	i = ix + iy*blockside
	return
}
func (s *claa) data(q *canvas.Path) {
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
func (s *claa) splitsv(np0, np1 mgl32.Vec2) {
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
func (s *claa) splitsh(np0, np1 mgl32.Vec2) {

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
func (s *claa) line(p0, p1 mgl32.Vec2) {
	cx, cy := int((p0[0]+p1[0])/2), int((p0[1]+p1[1])/2)
	b, i := s.offset(cx, cy)
	if b == -1 {
		return
	}
	//
	cover := int32((p1[1] - p0[1]) * float32(s.prc))
	len0 := int32((p0[0] - float32(cx)) * float32(s.prc))
	len1 := int32((p1[0] - float32(cx)) * float32(s.prc))
	area := int32((len0 + len1) * cover)
	//
	atomic.AddInt32(&s.buffer[b][i].area, area)
	atomic.AddInt32(&s.buffer[b][i].cover, cover)
}
func (s *claa) effective(c, a int32) uint32 {
	return s.toRange(s.iprc*c - a/2)
}
func (s *claa) toRange(value int32) uint32 {
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
