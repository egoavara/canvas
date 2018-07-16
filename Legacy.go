package canvas

//
////deprecated
//type Legacy struct {
//	buffer     *image.RGBA
//	swapbuffer *image.RGBA
//}
//
//func (s *Legacy) Type() SurfaceType {
//	return SurfaceTypeSoftware
//}
//func (s *Legacy) Bounds() image.Rectangle {
//	return s.buffer.Rect
//}
//func (s *Legacy) Clear() error {
//	for i := range s.swapbuffer.Pix {
//		s.swapbuffer.Pix[i] = 0
//	}
//	return nil
//}
//func (s *Legacy) Query(query *Path, shader Shader, transform *Transform) error {
//	if transform == nil {
//		transform = NewTransform()
//	}
//	// Stencil
//	var wg = new(sync.WaitGroup)
//	wg.Add(2)
//	go func() {
//		buf0 := make([]sortedSignIgnoreF32List, s.buffer.Rect.Max.Y)
//		if transform == nil || *transform == *NewTransform() {
//			//noThransformHorizontal(query, buf0)
//			withThransformHorizontal(query, transform, buf0)
//		} else {
//			withThransformHorizontal(query, transform, buf0)
//		}
//		s.hQ(buf0)
//		wg.Done()
//	}()
//	go func() {
//		buf1 := make([]sortedSignIgnoreF32List, s.buffer.Rect.Max.X)
//		if transform == nil || *transform == *NewTransform() {
//			//noThransformVertical(query, buf1)
//			withThransformVertical(query, transform, buf1)
//		} else {
//			withThransformVertical(query, transform, buf1)
//		}
//		s.vQ(buf1)
//		wg.Done()
//	}()
//	wg.Wait()
//	//
//	bound := s.mixQ()
//	if bound == invalidRect {
//		// no data do fill
//		return nil
//	}
//	// Fill
//	switch shd := shader.(type) {
//	case *ColorShader:
//		r, g, b, a := uint32(shd.R), uint32(shd.G), uint32(shd.B), uint32(shd.A)
//		wg.Add(bound.Dy() + 1)
//		for y := bound.Min.Y; y <= bound.Max.Y; y++ {
//			go func(y int) {
//				for x := bound.Min.X; x <= bound.Max.X; x++ {
//					offset := s.swapbuffer.PixOffset(x, y)
//					intensity := uint32(s.swapbuffer.Pix[offset+3])
//					s.swapbuffer.Pix[offset+3] = uint8(a * intensity / math.MaxUint8)
//					s.swapbuffer.Pix[offset+0] = uint8(r * intensity / math.MaxUint8)
//					s.swapbuffer.Pix[offset+1] = uint8(g * intensity / math.MaxUint8)
//					s.swapbuffer.Pix[offset+2] = uint8(b * intensity / math.MaxUint8)
//					//s.swapbuffer.Pix[offset+3] = uint8(a * intensity / math.MaxUint8)
//					//s.swapbuffer.Pix[offset+0] = u8min(s.swapbuffer.Pix[offset+3], uint8(r * intensity / math.MaxUint8))
//					//s.swapbuffer.Pix[offset+1] = u8min(s.swapbuffer.Pix[offset+3], uint8(r * intensity / math.MaxUint8))
//					//s.swapbuffer.Pix[offset+2] = u8min(s.swapbuffer.Pix[offset+3], uint8(r * intensity / math.MaxUint8))
//				}
//				wg.Done()
//			}(y)
//		}
//		wg.Wait()
//	}
//	return nil
//}
//func (s *Legacy) Flush() error {
//	s.buffer, s.swapbuffer = s.swapbuffer, s.buffer
//	s.Clear()
//	return nil
//}
//func (s *Legacy) Draw(dst draw.Image, r image.Rectangle, sp image.Point) {
//	rect := r.Intersect(s.buffer.Bounds().Add(sp))
//	for y := rect.Min.Y; y < rect.Max.Y; y++ {
//		for x := rect.Min.X; x < rect.Max.X; x++ {
//			dst.Set(x, y, s.buffer.At(x-rect.Min.X, y-rect.Min.Y))
//		}
//	}
//}
//
//func (s *Legacy) hQ(buf []sortedSignIgnoreF32List) {
//	wg := new(sync.WaitGroup)
//
//	for y, line := range buf {
//		if len(line) <= 0 {
//			continue
//		}
//		wg.Add(1)
//		go func(y int, line []float32) {
//			var dir = 0
//			var prev, curr float32
//			prev = mgl32.Abs(line[0])
//			dir = sign(dir, line[0])
//			for i := 1; i < len(line); i++ {
//				curr = mgl32.Abs(line[i])
//				//
//				offset := s.swapbuffer.Stride*y + 4*int(prev)
//				//s.swapbuffer.Pix[offset+0] = uint8((float32(int(prev + 1)) - prev) * math.MaxUint8)
//
//				s.swapbuffer.Pix[offset+0] = uint8((uint16(s.swapbuffer.Pix[offset+0]) + uint16((float32(int(prev+1))-prev)*math.MaxUint8)) / 2)
//				for x := int(prev + 1); x < int(curr); x++ {
//					offset = s.swapbuffer.Stride*y + 4*x
//					s.swapbuffer.Pix[offset+0] = math.MaxUint8
//				}
//				offset = s.swapbuffer.Stride*y + 4*int(curr)
//				s.swapbuffer.Pix[offset+0] = uint8((uint16(s.swapbuffer.Pix[offset+0]) + uint16((curr-float32(int(curr)))*math.MaxUint8)) / 2)
//				//
//				prev = curr
//				dir = sign(dir, line[i])
//			}
//			wg.Done()
//		}(y, line)
//	}
//	wg.Wait()
//}
//func (s *Legacy) vQ(buf []sortedSignIgnoreF32List) {
//	wg := new(sync.WaitGroup)
//
//	for x, line := range buf {
//		if len(line) <= 0 {
//			continue
//		}
//		wg.Add(1)
//		go func(x int, line []float32) {
//			var dir = 0
//			var prev, curr float32
//			prev = mgl32.Abs(line[0])
//			dir = sign(dir, line[0])
//			for i := 1; i < len(line); i++ {
//				curr = mgl32.Abs(line[i])
//				//
//				offset := s.swapbuffer.Stride*int(prev) + 4*x
//
//				s.swapbuffer.Pix[offset+1] = uint8((uint16(s.swapbuffer.Pix[offset+1]) + uint16((float32(int(prev+1))-prev)*math.MaxUint8)) / 2)
//				for y := int(prev + 1); y < int(curr); y++ {
//					offset = s.swapbuffer.Stride*y + 4*x
//					s.swapbuffer.Pix[offset+1] = math.MaxUint8
//				}
//				offset = s.swapbuffer.Stride*int(curr) + 4*x
//				s.swapbuffer.Pix[offset+1] = uint8((uint16(s.swapbuffer.Pix[offset+1]) + uint16((curr-float32(int(curr)))*math.MaxUint8)) / 2)
//
//				//
//				prev = curr
//				dir = sign(dir, line[i])
//			}
//			wg.Done()
//		}(x, line)
//	}
//	wg.Wait()
//
//}
//func (s *Legacy) mixQ() (res image.Rectangle) {
//	const threshold = 0.001
//	res = invalidRect
//
//	for y := 0; y < s.swapbuffer.Rect.Max.Y; y++ {
//		for x := 0; x < s.swapbuffer.Rect.Max.X; x++ {
//			offset := s.swapbuffer.PixOffset(x, y)
//			value := modify((float32(s.swapbuffer.Pix[offset+0])/math.MaxUint8 + float32(s.swapbuffer.Pix[offset+1])/math.MaxUint8) / 2)
//			if value <= threshold {
//				continue
//			}
//			res.Min.X = imin(res.Min.X, x)
//			res.Min.Y = imin(res.Min.Y, y)
//			res.Max.X = imax(res.Max.X, x)
//			res.Max.Y = imax(res.Max.Y, y)
//
//			s.swapbuffer.Pix[offset+0] = 0
//			s.swapbuffer.Pix[offset+1] = 0
//			s.swapbuffer.Pix[offset+3] = uint8((value) * math.MaxUint8)
//		}
//	}
//	return
//}
//
//
//func withThransformHorizontal(query *Path, transfrom *Transform, result []sortedSignIgnoreF32List) {
//	var l0, l1 mgl32.Vec3
//	var from, to mgl32.Vec3
//	var dir, delta float32
//	//
//	l0 = transfrom.rawMul(query.Data[0])
//	for i := 1; i < len(query.Data); i++ {
//		l1 = transfrom.rawMul(query.Data[i])
//		if math.IsNaN(float64(l0[0])) || math.IsNaN(float64(l1[0])) {
//			l0 = l1
//			continue
//		}
//		//
//		if l0[1] < l1[1] {
//			dir = 1
//			from = l0
//			to = l1
//		} else {
//			dir = -1
//			from = l1
//			to = l0
//		}
//		//
//		delta = (to[0] - from[0]) / (to[1] - from[1])
//		var temp = from[0]
//		for y := int(from[1]); y < int(to[1]); y++ {
//			result[y].Append(f32ZeroMula(temp, dir))
//			temp += delta
//		}
//		//
//		l0 = l1
//	}
//}
//func withThransformVertical(query *Path, transfrom *Transform, result []sortedSignIgnoreF32List) {
//	var l0, l1 mgl32.Vec3
//	var from, to mgl32.Vec3
//	var dir, delta float32
//	//
//	l0 = transfrom.rawMul(query.Data[0])
//	for i := 1; i < len(query.Data); i++ {
//		l1 = transfrom.rawMul(query.Data[i])
//		if math.IsNaN(float64(l0[0])) || math.IsNaN(float64(l1[0])) {
//			l0 = l1
//			continue
//		}
//		//
//
//		if l0[0] < l1[0] {
//			dir = 1
//			from = l0
//			to = l1
//		} else {
//			dir = -1
//			from = l1
//			to = l0
//		}
//		//
//		delta = (to[1] - from[1]) / (to[0] - from[0])
//		var temp = from[1]
//		for x := int(from[0]); x < int(to[0]); x++ {
//			result[x].Append(f32ZeroMula(temp, dir))
//			temp += delta
//		}
//		//
//		l0 = l1
//	}
//}
//
////func noThransformHorizontal(query *Path, result []sortedSignIgnoreF32List) {
////	var l0, l1 mgl32.Vec3
////	var from, to mgl32.Vec3
////	var dir, delta float32
////	//
////	l0 = query.Data[0]
////	for i := 1; i < len(query.Data); i++ {
////		l1 = query.Data[i]
////		if math.IsNaN(float64(l0[0])) || math.IsNaN(float64(l1[0])) {
////			l0 = l1
////			continue
////		}
////		//
////		if l0[1] < l1[1] {
////			dir = 1
////			from = l0
////			to = l1
////		} else {
////			dir = -1
////			from = l1
////			to = l0
////		}
////		//
////		delta = (to[0] - from[0]) / (to[1] - from[1])
////		var temp = from[0]
////		for y := int(from[1] + 0.9999); y < int(to[1]+.5); y++ {
////			result[y].Append(floatSign(temp, dir))
////			temp += delta
////		}
////		//
////		l0 = l1
////	}
////}
////func noThransformVertical(query *Path, result []sortedSignIgnoreF32List) {
////	var l0, l1 mgl32.Vec3
////	var from, to mgl32.Vec3
////	var dir, delta float32
////	//
////	l0 = query.Data[0]
////	for i := 1; i < len(query.Data); i++ {
////		l1 = query.Data[i]
////		if math.IsNaN(float64(l0[0])) || math.IsNaN(float64(l1[0])) {
////			l0 = l1
////			continue
////		}
////		//
////
////		if l0[0] < l1[0] {
////			dir = 1
////			from = l0
////			to = l1
////		} else {
////			dir = -1
////			from = l1
////			to = l0
////		}
////		//
////		delta = (to[1] - from[1]) / (to[0] - from[0])
////		var temp = from[1]
////		for x := int(from[0] + .9999); x < int(to[0] +.5); x++ {
////			result[x].Append(floatSign(temp, dir))
////			temp += delta
////		}
////		//
////		l0 = l1
////	}
////}
//func modify(a float32) float32 {
//	return -(a-1)*(a-1) + 1
//}
//
//func f32ZeroMula(a, b float32) float32 {
//	if a == 0 && b < 0{
//		return MinusZero
//	}
//	return a * b
//}