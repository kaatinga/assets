package assets

import (
	"github.com/kaatinga/strconv"
)

// String2Uint32 checks and converts input string to uint32 type.
func String2Uint32(input string) (uint32, error) {
	return faststrconv.GetUint32(input)
}

// Bytes2Uint32 checks and converts input string as []byte to uint32 type.
func Bytes2Uint32(input []byte) (uint32, error) {
	return faststrconv.GetUint32(input)
}

// String2Uint16 checks and converts input string to uint16 type.
func String2Uint16(input string) (uint16, error) {
	return faststrconv.GetUint16(input)
}

// Bytes2Uint16 checks and converts input string as []byte to uint16 type.
func Bytes2Uint16(input []byte) (uint16, error) {
	return faststrconv.GetUint16(input)
}

// String2Byte checks and converts input string to byte type.
func String2Byte(input string) (byte, error) {
	return faststrconv.GetByte(input)
}

// Bytes2Byte checks and converts input string as []byte to byte type.
func Bytes2Byte(input []byte) (byte, error) {
	return faststrconv.GetByte(input)
}

// Uint162Bytes converts an uint16 number to string.
func Uint162Bytes(num uint16) []byte {
	return faststrconv.Uint162Bytes(num)
}

// Uint162String converts an uint16 number to string.
func Uint162String(num uint16) string {
	return faststrconv.Uint162String(num)
}

// Byte2Bytes converts a byte number to []byte.
func Byte2Bytes(num byte) []byte {
	return faststrconv.Byte2Bytes(num)
}

// Byte2String converts a byte number to string.
func Byte2String(num byte) string {
	return faststrconv.Byte2String(num)
}
