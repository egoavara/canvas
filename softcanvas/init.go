package softcanvas

import "github.com/iamGreedy/canvas"
const Driver = canvas.Driver("software")

func init() {
	canvas.Setup(Driver, func(w, h int, options ...canvas.Option) (canvas.Surface, error) {
		s := NewSoftware(w,h, options...)
		return s, nil
	})
}
