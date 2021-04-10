package assets

const (
	UnicodeMaskUint32 uint32 = 0xf
	UnicodeMaskByte   byte   = 0xf
	ByteLengthMask    int    = 0x3
)

// String2Uint32 checks and converts input string to uint32 type.
func String2Uint32(input string) (output uint32, err error) {

	if len(input) == 0 {
		return 0, ErrNotUint32
	}

	if len(input) > 10 {
		return 0, ErrNumberExceedMaxUint32Value
	}

	var i int
	for {

		if i == len(input) {
			break
		}

		if input[i] < 0x30 || input[i] > 0x39 {
			return 0, ErrNotUint32
		}

		output = (output << 3) + (output << 1) + uint32(input[i])&UnicodeMaskUint32

		if output == 0 && i > 8 {
			return 0, ErrNumberExceedMaxUint32Value
		}

		i++
	}

	return
}

// Bytes2Uint32 checks and converts input string as []byte to uint32 type.
func Bytes2Uint32(input []byte) (output uint32, err error) {

	if len(input) == 0 {
		return 0, ErrNotUint32
	}

	if len(input) > 10 {
		return 0, ErrNumberExceedMaxUint32Value
	}

	var i int
	for {

		if i == len(input) {
			break
		}

		if input[i] < 0x30 || input[i] > 0x39 {
			return 0, ErrNotUint32
		}

		output = (output << 3) + (output << 1) + uint32(input[i])&UnicodeMaskUint32

		if output == 0 && i > 8 {
			return 0, ErrNumberExceedMaxUint32Value
		}

		i++
	}

	return
}

// String2Byte checks and converts input string to uint32 type.
func String2Byte(input string) (output byte, err error) {

	length := len(input)
	if length&^ByteLengthMask != 0 {
		return 0, ErrNotByte
	}

	var i int
	for {

		if i == length {
			break
		}

		if input[i] < 0x30 || input[i] > 0x39 {
			return 0, ErrNotByte
		}

		output = (output << 3) + (output << 1) + input[i]&UnicodeMaskByte

		if output == 0 && i > 1 {
			return 0, ErrNumberExceedMaxByteValue
		}

		i++
	}

	return output, nil
}
