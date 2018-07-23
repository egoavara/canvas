package canvas

import (
	"os"
	"image/png"
	"github.com/pkg/errors"
)

func Capture(path string, surface Surface) (err error) {
	var f *os.File
	f, err = os.OpenFile(path, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer func() {
		f.Close()
		if err != nil {
			os.Remove(path)
		}
	}()
	//
	var fb FrameBuffer
	if IsFail(surface.Option(&fb)){
		return errors.New("Framebuffer fail to get")
	}
	//
	return png.Encode(f, fb.Convert())
}
