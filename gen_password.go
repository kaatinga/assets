package assets

import (
	"math/rand"
)

const (
	chars       = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	charsLength = len(chars)
)

// GenPassword generates a password of a set length.
func GenPassword(length byte) string {
	return string(GenPasswordAsBytes(length))
}

// GenPasswordAsBytes generates a password of a set length.
func GenPasswordAsBytes(length byte) []byte {

	if length == 0 { // in case we do not want to point out the length we can set zero
		length = 7
	}

	password := make([]byte, length)

	for i := byte(0); i < length; i++ {
		password[i] = chars[rand.Intn(charsLength)]
	}

	return password
}
