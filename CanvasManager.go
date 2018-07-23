package canvas


// map[drivername] driver
var manager = make(map[Driver]func(w, h int, options ... Option) (Surface, error))
var first *Driver
func Setup(name Driver, driver func(w, h int, options ... Option) (Surface, error)) {
	if _, ok := manager[name]; ok{
		panic("Driver name ")
	}
	if first == nil{
		first = &name
	}
	manager[name] = driver
}

func NewSurface(driver Driver, width, height int, options ... Option) (Surface, error) {
	return manager[driver](width, height, options...)
}