package assets

// IntError - Error type based on int value.
type IntError int

// Error returns error description.
func (ie IntError) Error() string {
	return errorDescriptions[ie]
}

// errorDescriptions contains descriptions for the IntError errors.
var errorDescriptions = []string{
	ErrNotMoreThan255ValuesAreSupported: "the number of input values exceeds 255",
	ErrAtLeast2ValuesNeeded:             "the number of input values is less than 2",
}

const (
	ErrNotMoreThan255ValuesAreSupported IntError = 6 + iota
	ErrAtLeast2ValuesNeeded
)
