package assets

import (
	"strconv"
)

// StUint16 converts input string to uint16 type
func StUint16(inputString string) (result uint16, ok bool) {

	var err error // to store an error if it occurs
	var tmpUint16 uint64 // a temporary uint64 value

	tmpUint16, err = strconv.ParseUint(inputString, 10, 64)
	result = uint16(tmpUint16)
	if err == nil {
		ok = true
	}

	return
}

// StBool converts input string to bool type
func StBool(inputString string) (result bool, ok bool) {

	var err error // to store an error if it occurs

	result, err = strconv.ParseBool(inputString)
	if err == nil {
		ok = true
	}

	return
}