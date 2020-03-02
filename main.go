package assets

import (
	"strconv"
)

// StUint16 converts input string to uint16 type
func StUint16(inputString string) (result uint16, ok bool) {

	var err error // to store error result
	var tmpUint16 uint64 // a temporary uint64 value

	tmpUint16, err = strconv.ParseUint(inputString, 10, 64)
	result = uint16(tmpUint16)
	if err == nil {
		ok = true
	}

	return
}

// StByte converts input string to Byte type
func StByte(inputString string) (result byte, ok bool) {

	var err error       // to store error result
	var tmpInt64 int // a temporary uint64 value

	tmpInt64, err = strconv.Atoi(inputString)
	if err == nil {
		ok = true
		result = byte(tmpInt64)
	}

	return
}

// StBool converts input string "true" to bool type
func StBool(inputString string) bool {

	if inputString == "true" || inputString == "on" {
		return true
	}

	return false
}