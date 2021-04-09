package assets

const UnicodeMask = 0xf

// String2Uint32 checks and converts input string to uint32 type.
func String2Uint32(input string) (output uint32, err error) {

	if len(input) == 0 {
		return 0, ErrNotUint32
	}

	var i int
	for {

		if i == len(input) {
			break
		}

		if input[i] < 0x30 || input[i] > 0x39 {
			return 0, ErrNotUint32
		}

		output = (output << 3) + (output << 1) + uint32(input[i])&UnicodeMask

		if output == 0 && i > 8 {
			return 0, ErrNumberExceedMaxUint32Value
		}

		i++
	}

	return
}
