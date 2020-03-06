package assets

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// StUint16 converts input string to uint16 type
func StUint16(inputString string) (uint16, bool) {
	var err error  // to store error result
	var tmpInt int // a temporary int value

	tmpInt, err = strconv.Atoi(inputString)
	if err == nil {
		if tmpInt >= 0 && tmpInt < 65536 {
			return uint16(tmpInt), true
		}
	}
	return 0, false
}

// StByte converts input string to Byte type
func StByte(inputString string) (byte, bool) {
	var err error  // to store error result
	var tmpInt int // a temporary int value

	tmpInt, err = strconv.Atoi(inputString)
	if err == nil {
		if tmpInt >= 0 && tmpInt < 256 {
			return byte(tmpInt), true
		}
	}
	return 0, false
}

// StBool converts input string "true" to bool type
func StBool(inputString string) bool {

	if inputString == "true" || inputString == "on" {
		return true
	}

	return false
}

// GenPassword generates a password of a set length
func GenPassword(length byte) (str string, err error) {
	rand.Seed(time.Now().UnixNano()) // in order to issue really random password
	chars := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	charsLength := len(chars)

	if length == 0 { // in case we do not want to point out the length we can set zero
		length = 7
	}

	var builder strings.Builder

	var i byte
	for ; i < length; i++ {
		err = builder.WriteByte(chars[rand.Intn(charsLength)])
		if err == nil {
			return builder.String(), nil
		}
	}
	return "", err
}