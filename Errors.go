package canvas

import "github.com/pkg/errors"

// Predefined Errors for implement Surface.
// They are not used in the package, but are designed for surface implementation in other packages.
//
var (
	// Surface
	ErrorUnknownDriver = errors.New("Unknown driver")
	ErrorAllocation    = errors.New("Allocation fail")
	ErrorNotAllowNil   = errors.New("nil is not allow")
	ErrorTooBusy       = errors.New("Currently, Query queue is too busy to handle this")
	ErrorUnsupported   = errors.New("Unsupported Option")
	ErrorClosed        = errors.New("This is closed Surface, Do not use this")
	//
)

// Fetal is a very serious error.
// They are passed through the panic(), not the return value.
var (
	FetalSetup = errors.New("Fetal error, Setup check fail")
)
