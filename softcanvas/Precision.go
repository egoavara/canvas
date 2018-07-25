package softcanvas

import "github.com/iamGreedy/canvas"

type Precision float32

const (
	X1  Precision = 1
	X2  Precision = 2
	X4  Precision = 4
	X8  Precision = 8
	X16 Precision = 16
)



func (Precision) OptionType() canvas.OptionType {
	return canvas.Write | canvas.Read | canvas.Init
}