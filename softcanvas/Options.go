package softcanvas

import (
	"github.com/iamGreedy/canvas/claa"
	"github.com/iamGreedy/canvas"
)

type (
	OptionCLAAPrecision claa.Precision
)

func (OptionCLAAPrecision) OptionType() canvas.OptionType {
	return canvas.Write | canvas.Read | canvas.Init
}

