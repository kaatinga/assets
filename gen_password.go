package assets

import (
	"math/rand"
	"strings"
)

const (
	chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	charsLength = len(chars)
)

// GenPassword generates a password of a set length.
func GenPassword(length byte) (string, error) {

	if length == 0 { // in case we do not want to point out the length we can set zero
		length = 7
	}

	var (
		builder strings.Builder
		err     error
	)

	for i := byte(0); i < length; i++ {
		err = builder.WriteByte(chars[rand.Intn(charsLength)]) //nolint:gosec
		if err != nil {
			return "", err
		}
	}

	return builder.String(), nil
}
