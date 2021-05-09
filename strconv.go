package assets

const (
	UnicodeMaskUint64 uint64 = 0xf
	UnicodeMaskUint32 uint32 = 0xf
	UnicodeMaskUint16 uint16 = 0xf

	ByteLengthMask int = 0b11

	Uint16LengthMask int = 0b101

	Uint32LengthMask1 int = 0b111  // checks the number > 7
	Uint32LengthMask2 int = 0b1010 // checks the number == 8, 10
	Uint32LengthMask3 int = 0b1001 // checks the number == 9
)

// String2Uint32 checks and converts input string to uint32 type.
func String2Uint32(input string) (uint32, error) {

	if !(len(input)&^Uint32LengthMask1 == 0 ||
		len(input)&^Uint32LengthMask2 == 0 ||
		len(input) == Uint32LengthMask3) ||
		len(input) == 0 {
		return 0, ErrNotUint32
	}

	var i int
	var output uint64
	for {

		if input[i] < 0x30 || input[i] > 0x39 {
			return 0, ErrNotUint32
		}

		output = (output << 3) + (output << 1) + uint64(input[i])&UnicodeMaskUint64

		if output&^0xffffffff != 0 {
			return 0, ErrNumberExceedMaxUint32Value
		}

		i++

		if i == len(input) {
			break
		}
	}

	return uint32(output), nil
}

// Bytes2Uint32 checks and converts input string as []byte to uint32 type.
func Bytes2Uint32(input []byte) (uint32, error) {

	if !(len(input)&^Uint32LengthMask1 == 0 ||
		len(input)&^Uint32LengthMask2 == 0 ||
		len(input) == Uint32LengthMask3) ||
		len(input) == 0 {
		return 0, ErrNotUint32
	}

	var i int
	var output uint64
	for {

		if input[i] < 0x30 || input[i] > 0x39 {
			return 0, ErrNotUint32
		}

		output = (output << 3) + (output << 1) + uint64(input[i])&UnicodeMaskUint64

		if output&^0xffffffff != 0 {
			return 0, ErrNumberExceedMaxUint32Value
		}

		i++

		if i == len(input) {
			break
		}
	}

	return uint32(output), nil
}

// String2Uint16 checks and converts input string to uint16 type.
func String2Uint16(input string) (uint16, error) {

	if !(len(input)&^Uint16LengthMask == 0 ||
		len(input)&^ByteLengthMask == 0) {
		return 0, ErrNotUint16
	}

	var i int
	var output uint32
	for {

		if input[i] < 0x30 || input[i] > 0x39 {
			return 0, ErrNotUint16
		}

		output = (output << 3) + (output << 1) + uint32(input[i])&UnicodeMaskUint32

		if output&^0xffff != 0 {
			return 0, ErrNumberExceedMaxUint16Value
		}

		i++

		if i == len(input) {
			break
		}
	}

	return uint16(output), nil
}

// Bytes2Uint16 checks and converts input string as []byte to uint16 type.
func Bytes2Uint16(input []byte) (uint16, error) {

	if !(len(input)&^Uint16LengthMask == 0 ||
		len(input)&^ByteLengthMask == 0) {
		return 0, ErrNotUint16
	}

	var i int
	var output uint32
	for {

		if input[i] < 0x30 || input[i] > 0x39 {
			return 0, ErrNotUint16
		}

		output = (output << 3) + (output << 1) + uint32(input[i])&UnicodeMaskUint32

		if output&^0xffff != 0 {
			return 0, ErrNumberExceedMaxUint16Value
		}

		i++

		if i == len(input) {
			break
		}
	}

	return uint16(output), nil
}

// String2Byte checks and converts input string to uint32 type.
func String2Byte(input string) (byte, error) {

	if len(input)&^ByteLengthMask != 0 || len(input) == 0 {
		return 0, ErrNotByte
	}

	var i int
	var output uint16
	for {
		if input[i] < 0x30 || input[i] > 0x39 {
			return 0, ErrNotByte
		}

		output = uint16(input[i])&UnicodeMaskUint16 + (output << 3) + (output << 1)

		if output&^0xff != 0 {
			return 0, ErrNumberExceedMaxByteValue
		}

		i++

		if i == len(input) {
			break
		}
	}

	return byte(output), nil
}

// Bytes2Byte checks and converts input string as []byte to uint32 type.
func Bytes2Byte(input []byte) (byte, error) {

	if len(input)&^ByteLengthMask != 0 || len(input) == 0 {
		return 0, ErrNotByte
	}

	var i int
	var output uint16
	for {
		if input[i] < 0x30 || input[i] > 0x39 {
			return 0, ErrNotByte
		}

		output = uint16(input[i])&UnicodeMaskUint16 + (output << 3) + (output << 1)

		if output&^0xff != 0 {
			return 0, ErrNumberExceedMaxByteValue
		}

		i++

		if i == len(input) {
			break
		}
	}

	return byte(output), nil
}
