package assets

// IntError - Error type based on int value.
type IntError int

// Error returns error description.
func (err IntError) Error() string {
	return errorDescriptions[err]
}

// errorDescriptions contains descriptions for the IntError errors.
var errorDescriptions = []string{
	0: "the input string is not an uint32 number",
	1: "the input string is a number, but the value exceeds the maximum value",
	2: "the input string is not an uint8 number",
	3: "the input string is a number, but the value exceeds the maximum value",
}

const (

	// Uint32 conversion errors
	ErrNotUint32 IntError = iota
	ErrNumberExceedMaxUint32Value

	// Byte conversion errors
	ErrNotByte
	ErrNumberExceedMaxByteValue
)
