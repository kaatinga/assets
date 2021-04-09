package assets

// IntError - Error type based on int value.
type IntError int

// Error returns error description.
func (err IntError) Error() string {
	return errorDescriptions[err]
}

// errorDescriptions contains descriptions for the IntError errors.
var errorDescriptions = []string{
	0: "the input string is not a number",
	1: "the input string is a number, but the value exceeds the maximum value",
}

const (

	// String2Uint32 errors:
	ErrNotUint32 IntError = iota
	ErrNumberExceedMaxUint32Value
)
