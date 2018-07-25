package canvas

// Buffer support option
type (
	// Use Flush buffer for performance
	//
	// = Send
	//
	UseQueryBuffer struct{}

	// GetCurrent Buffer Query Count
	//
	// = Read
	//
	QueryBufferLength int

	// Close buffer
	//
	// = Send
	CloseQueryBuffer struct{}

	// Flush all buffer for Data
	//
	// = Send
	FlushQueryBuffer struct{}

	// Lock until every flushing success
	//
	// = Send
	WaitFlush struct{}
)

func (UseQueryBuffer) OptionType() OptionType {
	return Send
}
func (QueryBufferLength) OptionType() OptionType {
	return Read
}
func (CloseQueryBuffer) OptionType() OptionType {
	return Send
}
func (FlushQueryBuffer) OptionType() OptionType {
	return Send
}
func (WaitFlush) OptionType() OptionType {
	return Send
}