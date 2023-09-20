package assets

import (
	"unicode"
)

const allowedCharactersCyrillicString = `&"+-»«.,/№`

var symbolRange = []*unicode.RangeTable{
	unicode.Cyrillic,
	unicode.Digit,
	unicode.Space,
}

// IsCyrillicString check if an only allowed set of symbols is in the company name.
func IsCyrillicString(company string) bool {
	var found bool
	for _, value := range company {
		found = unicode.IsOneOf(symbolRange, value)
		if found {
			continue
		}

		for _, symbol := range allowedCharactersCyrillicString {
			if value == symbol {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	return true
}

// CheckName check if an only allowed set of symbols is in the string.
func CheckName(name string) bool {
	for _, value := range name {
		if !unicode.IsOneOf([]*unicode.RangeTable{
			unicode.Cyrillic,
			unicode.Space,
		}, value) {
			return false
		}
	}

	return true
}
