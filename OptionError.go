package canvas


// There is two type Fail, one is .Option(...) return error wrap by ResultFail
// Other is return nil, which is so obious, no need to explain
//
type (
	ResultFail struct {
		Cause error
	}
)
func (ResultFail) OptionType() OptionType {
	return Error
}
func (s ResultFail) Error() string {
	return s.Cause.Error()
}

func IsFail(o Option) bool {
	if o == nil{
		return true
	}
	return o.OptionType() & Error == Error
}

