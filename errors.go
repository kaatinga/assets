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
	1: "the input string is a number, but the value exceeds the maximum uint32 value",
	2: "the input string is not an uint16 number",
	3: "the input string is a number, but the value exceeds the maximum uint16 value",
	4: "the input string is not an uint8 number",
	5: "the input string is a number, but the value exceeds the maximum uint8 value",
	6: "the number of input values exceeds 255",
	7: "the number of input values is less than 2",
}

const (

	// Uint32 conversion errors
	ErrNotUint32 IntError = iota
	ErrNumberExceedMaxUint32Value

	// Uint16 conversion errors
	ErrNotUint16
	ErrNumberExceedMaxUint16Value

	// Byte conversion errors
	ErrNotByte
	ErrNumberExceedMaxByteValue

	// Bool comparsion
	ErrNotMoreThan255ValuesAreSupported
	ErrAtLeast2ValuesNeeded
)
