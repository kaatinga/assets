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

// String2Byte checks and converts input string to byte type.
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

// Bytes2Byte checks and converts input string as []byte to byte type.
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

const (
	uint1610    uint16 = 10
	uint16100   uint16 = 100
	uint161000  uint16 = 1000
	uint1610000 uint16 = 10000
)

// Uint162Bytes converts an uint16 number to string.
func Uint162Bytes(num uint16) []byte {

	convertedNumber, i := getSliceUint16(num)

	for {
		convertedNumber[i] = byte(num%10) | 0x30
		num = num / 10
		if i == 0 {
			return convertedNumber
		}
		i--
	}
}

// Uint162String converts an uint16 number to string.
func Uint162String(num uint16) string {
	return string(Uint162Bytes(num))
}

func getSliceUint16(num uint16) ([]byte, int) {
	if num < uint1610 {
		return make([]byte, 1), 0
	}

	if num < uint16100 {
		return make([]byte, 2), 1
	}

	if num < uint161000 {
		return make([]byte, 3), 2
	}

	if num < uint1610000 {
		return make([]byte, 4), 3
	}

	return make([]byte, 5), 4
}

//// Uint322Bytes converts an uint32 number to string.
//func Uint322Bytes(num uint32) []byte {
//
//	convertedNumber, i := getSliceUint32(num)
//
//	for {
//		convertedNumber[i] = byte(num%10) | 0x30
//		num = num / 10
//		if i == 0 {
//			return convertedNumber
//		}
//		i--
//	}
//}
//
//const (
//	uint3210         uint32 = 10
//	uint32100        uint32 = 100
//	uint321000       uint32 = 1000
//	uint3210000      uint32 = 10000
//	uint32100000     uint32 = 100000
//	uint321000000    uint32 = 1000000
//	uint3210000000   uint32 = 10000000
//	uint32100000000  uint32 = 100000000
//	uint321000000000 uint32 = 1000000000
//)
//
//func getSliceUint32(num uint32) ([]byte, int) {
//
//	if num < uint3210 {
//		return make([]byte, 1), 0
//	}
//
//	if num < uint32100 {
//		return make([]byte, 2), 1
//	}
//
//	if num < uint321000 {
//		return make([]byte, 3), 2
//	}
//
//	if num < uint3210000 {
//		return make([]byte, 4), 3
//	}
//
//	if num < uint32100000 {
//		return make([]byte, 5), 4
//	}
//
//	if num < uint321000000 {
//		return make([]byte, 6), 5
//	}
//
//	if num < uint3210000000 {
//		return make([]byte, 7), 6
//	}
//
//	if num < uint32100000000 {
//		return make([]byte, 8), 7
//	}
//
//	if num < uint321000000000 {
//		return make([]byte, 9), 8
//	}
//
//	return make([]byte, 10), 9
//}

//func getSize(num uint16, size int) int {
//	if size == 5 {
//		return size
//	}
//	if num < uint16sizes[size] {
//		return size
//	}
//	return getSize(num, size+1)
//}
//
//func convertNumber(convertedNumber []byte, i int, num uint16) {
//	convertedNumber[i] = byte(num%10) + 0x30
//	num = num / 10
//	if num == 0 {
//		return
//	}
//	convertNumber(convertedNumber, i-1, num)
//}

// Byte2Bytes converts a byte number to []byte.
func Byte2Bytes(num byte) []byte {
	convertedNumber, i := getSliceByte(num)

	for {
		convertedNumber[i] = num%10 | 0x30
		num = num / 10
		if i == 0 {
			return convertedNumber
		}
		i--
	}
}

// Byte2String converts a byte number to string.
func Byte2String(num byte) string {
	return string(Byte2Bytes(num))
}

const (
	byte10  byte = 10
	byte100 byte = 100
)

func getSliceByte(num byte) ([]byte, int) {
	if num < byte10 {
		return make([]byte, 1), 0
	}

	if num < byte100 {
		return make([]byte, 2), 1
	}

	return make([]byte, 3), 2
}
