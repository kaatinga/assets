package assets

import (
	"strings"
	"unicode"
)

const allowedCharactersRussianCompanyName = "& \"+-»«№"

var symbolRange = []*unicode.RangeTable{
	unicode.Cyrillic,
	unicode.Digit,
}

// CheckRussianCompanyName check if an only allowed set of symbols is in the company name.
func CheckRussianCompanyName(company string) bool {
	company = RemoveCharacters(company, allowedCharactersRussianCompanyName) // to remove a set of allowed symbols

	for _, value := range company {
		if !unicode.IsOneOf(symbolRange, value) {
			return false
		}
	}

	return true
}

// CheckName check if an only allowed set of symbols is in the string.
func CheckName(name string) bool {
	name = RemoveCharacters(name, " ") // to remove space

	for _, value := range name {
		if !unicode.In(value, unicode.Cyrillic) {
			return false
		}
	}

	return true
}

// RemoveCharacters removes the set of characters from the input string.
func RemoveCharacters(input, characters string) string {
	filter := func(r rune) rune {
		if !strings.ContainsRune(characters, r) {
			return r
		}
		return -1
	}
	return strings.Map(filter, input)
}
