package canvas

import (
	"os"
	"image/png"
	"image"
)

func Capture(path string, surface Surface) error {
	f, err := os.OpenFile(path, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	//
	rgba := image.NewRGBA(surface.Bounds())
	surface.Draw(rgba, rgba.Rect, image.ZP)
	//
	return png.Encode(f, rgba)
}
