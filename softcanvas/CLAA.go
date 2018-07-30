package softcanvas

import (
	"github.com/go-gl/mathgl/mgl32"
	"image"
	"math"
	"sync"
)

const blockside = 64
const blocksize = blockside * blockside

const (
	l1 = 1  // 1pixel per render
	l2 = 8  // 64pixel as one pix render
	l3 = 16 // 256pixel as on pixel render
)

type (
	claaManager struct {
		ref *Software
		//
		wait []*claablock
		lck  sync.Mutex
	}
	claa interface {
		data(points ...mgl32.Vec3)
		close()
		run(fn func(x, y int, effective uint32))
	}

	claablock [blocksize]cell

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
func (s *claaManager) get(rectangle image.Rectangle) claa {
	var needw, needh, left int
	dx, dy := rectangle.Dx(), rectangle.Dy()
	if needw, left = dx/blockside, dx%blockside; left != 0 {
		needw += 1
	}
	if needh, left = dy/blockside, dy%blockside; left != 0 {
		needh += 1
	}
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

		return &claa1layer{
			buf1w: needw,
			bound: rectangle,
			buf1:  temp,
			man:   s,
			prc:   float32(s.ref.precision),
			iprc:  int32(s.ref.precision),
		}
	}
	temp := s.wait[:totalneed]
	s.wait = s.wait[totalneed:]
	return &claa1layer{
		buf1w: needw,
		bound: rectangle,
		buf1:  temp,
		man:   s,
		prc:   float32(s.ref.precision),
		iprc:  int32(s.ref.precision),
	}
}
func (s *claaManager) max() int {
	return (s.ref.size.X/blockside + 1) * (s.ref.size.Y/blockside + 1)
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
