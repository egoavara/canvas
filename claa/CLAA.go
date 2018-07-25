package claa

import (
	"github.com/go-gl/mathgl/mgl32"
	"math"
	"sync"
)

const (
	X1  Precision = 1
	X2  Precision = 2
	X4  Precision = 4
	X8  Precision = 8
	X16 Precision = 16
)

type (
	CLAA struct {
		w, h   int
		p      Precision
		buffer []Cell
	}
	Cell struct {
		cover int32
		area  int32
	}
	Precision float32
	DataType  uint8
)

func NewCLAA(precision Precision, width, height int) *CLAA {
	return &CLAA{
		w:      width,
		h:      height,
		p:      precision,
		buffer: make([]Cell, width*height),
	}
}

// points are closed path
func (s *CLAA) Clear() {
	for i := range s.buffer {
		s.buffer[i].cover = 0
		s.buffer[i].area = 0
	}
}

const (
	RAW             DataType = iota
	ALIGNED         DataType = iota
)

func (s *CLAA) GetPrecision() Precision{
	return s.p
}
func (s *CLAA) GetWidth() int{
	return s.w
}
func (s *CLAA) GetHeight() int{
	return s.h
}
func (s *CLAA) SetPrecision(p Precision) {
	s.Clear()
	s.p =p
}
func (s *CLAA) Data(datatype DataType, points ...mgl32.Vec3) {
	switch datatype {
	case RAW:
		points = AlignGrid(s.p, points)
	case ALIGNED:
	}
	from := 0
	for to := 0; to < len(points); to++ {
		if math.IsNaN(float64(points[to][0])) {
			if to-from > 3 {
				lst := points[from:to]
				for i := 1; i < len(lst); i++ {
					l0, l1 := lst[i-1], lst[i]
					cx, cy := int((l0[0]+l1[0])/2), int((l0[1]+l1[1])/2)
					if cy >= s.h || cx >= s.w {
						continue
					}
					cover := int32((l1[1] - l0[1]) * float32(s.p))
					len0 := int32((l0[0] - float32(cx)) * float32(s.p))
					len1 := int32((l1[0] - float32(cx)) * float32(s.p))
					area := int32((len0 + len1) * cover)
					//
					offset := s.w*cy + cx
					s.buffer[offset].area += area
					s.buffer[offset].cover += cover
				}
			}
			to++
			from = to
			continue
		}
	}
}
func (s *CLAA) Result(dst []uint8, stride, pixsize, offset int) {
	wg := new(sync.WaitGroup)
	width := int32(s.p)
	wg.Add(s.h)
	//
	for y := 0; y < s.h; y++ {
		go func(y int) {
			var prev int32 = 0
			for x := 0; x < s.w; x++ {
				var res int32
				cell := s.buffer[s.w*y+x]
				res = effective(width, cell.cover+prev, cell.area)
				if cell.area != 0 || cell.cover != 0 {
					prev += cell.cover
				}
				if res == 0 {
					continue
				}

				offset := stride*y + pixsize*x + offset
				dst[offset] = toRange(res, width)
			}
			wg.Done()
		}(y)
	}
	wg.Wait()
}
func (s *CLAA) SyncResult(dst []uint8, stride, pixsize, offset int) {
	width := int32(s.p)
	for y := 0; y < s.h; y++ {
		var prev int32 = 0
		for x := 0; x < s.w; x++ {
			var res int32
			cell := s.buffer[s.w*y+x]
			res = effective(width, cell.cover+prev, cell.area)
			if cell.area != 0 || cell.cover != 0 {
				prev += cell.cover
			}
			if res == 0 {
				continue
			}
			offset := stride*y + pixsize*x + offset
			dst[offset] = toRange(res, width)
		}
	}
}

func effective(w, c, a int32) int32 {
	return w*c - a/2
}
func toRange(value int32, ip int32) uint8 {
	m := 16 / ip
	res := value * m * m
	if res < 0 {
		res = -res
	}
	if res > math.MaxUint8 {
		res = math.MaxUint8
	}
	return uint8(res)
}
