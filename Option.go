package canvas



// Option is an interface used for various synchronization, value transfer and acquisition on the surface.
// Each option allows you to pass in, receive, or synchronize values that match your name (structure name).
// All options except Essensial are optional implementations and do not necessarily exist.
// This can be checked with the .Support method of the Surface interface.
//
//
//
type Option interface {
	// They do not do anything, but they tell us that they are implementations of the Option interface.
	// It exists for a structure that can be utilized in IDE.
	OptionType() OptionType
}

// OptionType is the type that exists for the return of the .OptionType() method of the Option interface,
//  which tells you roughly what the implementation of the Option interface has.
type OptionType uint8


const (
	// Read 는
	Read  OptionType = 0x01
	Write OptionType = 0x02
	Init  OptionType = 0x04
	Send  OptionType = 0x08
	Error OptionType = 0x80
	//
	// Read  		OptionType
	// Write  		OptionType
	// Init  		OptionType
	// Send  		OptionType
	//
	// empty  		OptionType
	// empty 		OptionType
	// empty 		OptionType
	// ErrorFlag  	OptionType
)