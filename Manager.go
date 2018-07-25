package canvas

import (
	"github.com/pkg/errors"
	"fmt"
)

// map for drivers
var (
	manager = make(map[Driver]func(w, h int, options ... Option) (Surface, error))
	// First setup driver, "." is call this
	first *Driver
)

// Setup new driver for surface
// Once setup driver, that driver can be use by NewSurface
func Setup(name Driver, driver func(w, h int, options ... Option) (Surface, error)) {
	if _, ok := manager[name]; ok{
		panic(errors.WithMessage(FetalSetup, fmt.Sprintf("'%s' is already exist Driver name", string(name))))
	}
	if first == nil{
		first = &name
	}
	suf, err := driver(1,1)
	if err != nil {
		panic(errors.WithMessage(FetalSetup, fmt.Sprintf("This Driver is invalid, fail to make new Surface cause : '%s'", err.Error())))
	}
	if !suf.Support(checkvaliddriver...){
		panic(errors.WithMessage(FetalSetup, "This Driver is invalid, not support all essential drivers"))
	}
	// Setup driver success
	manager[name] = driver
}

// Make new Surface given driver
// If you don't care any driver
// set driver as "."
// "." mean any driver
func NewSurface(driver Driver, width, height int, options ... Option) (Surface, error) {
	if driver == "."{
		if first == nil{
			return nil, errors.WithMessage(ErrorUnknownDriver, "There is no available driver")
		}
		driver = *first
	}
	if fn, ok := manager[driver];ok{
		return fn(width, height, options...)
	}
	return nil, errors.WithMessage(ErrorUnknownDriver, fmt.Sprintf("'%s' is unknown", driver))
}