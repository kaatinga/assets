package assets

import (
	"encoding/json"
	"errors"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
	"unicode"
)

func init() {
	// in order to issue really random password
	rand.Seed(time.Now().UnixNano())
}

// String is an extended version of string type
type String struct {
	Parameter string
	Ok        bool
}

// SetStringByPointer checks and sets input Strings parameter to a string through a pointer
func (input *String) SetStringByPointer(output *string) bool {
	if (*input).Parameter != "" {
		*output = (*input).Parameter
		return true
	}
	return false
}

// StBool converts input string to bool type
func StBool(inputString string) bool {
	return inputString == "true" || inputString == "on"
}

// GenPassword generates a password of a set length
func GenPassword(length byte) (string, error) {
	chars := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	charsLength := len(chars)

	if length == 0 { // in case we do not want to point out the length we can set zero
		length = 7
	}

	var (
		builder strings.Builder
		err     error
	)

	var i byte
	for ; i < length; i++ {
		err = builder.WriteByte(chars[rand.Intn(charsLength)])
		if err != nil {
			return "", err
		}
	}
	return builder.String(), nil
}

// GetRandomByte generates a random byte number.
func GetRandomByte(max byte) byte {

	switch max {
	case 0:
		return 0
	default:
		return byte(rand.Int31n(int32(max)))
	}
}

// SaveFile saves a file in JSON format.
func SaveFile(data interface{}, path string) error {
	dataToWrite, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return os.WriteFile(path, dataToWrite, 0660)
}

// SafeQM escapes quotation marks adding '\' before them.
func SafeQM(str string) string {
	return strings.Replace(str, `"`, `\"`, -1)
}

// RemoveSafeQM removes symbols '\' before quotation marks.
func RemoveSafeQM(str string) string {
	return strings.Replace(str, `\"`, `"`, -1)
}

// CheckRussianCompanyName check if an only allowed set of symbols is in the company name
func CheckRussianCompanyName(company string) bool {

	// Russian company can have digits and russian symbols, as well as some symbols below
	var symbolRange = []*unicode.RangeTable{
		unicode.Cyrillic,
		unicode.Digit,
	}

	company = RemoveCharacters(company, "& \"+-»«") // to remove a set of allowed symbols

	for _, value := range company {
		if !unicode.IsOneOf(symbolRange, value) {
			return false
		}
	}

	return true
}

// CheckName check if an only allowed set of symbols is in the string
func CheckName(name string) bool {
	name = RemoveCharacters(name, " ") // to remove space

	for _, value := range name {
		if !unicode.In(value, unicode.Cyrillic) {
			return false
		}
	}

	return true
}

// RemoveCharacters removes the set of characters from the input string
func RemoveCharacters(input, characters string) string {
	filter := func(r rune) rune {
		if strings.IndexRune(characters, r) < 0 {
			return r
		}
		return -1
	}
	return strings.Map(filter, input)
}

// HTTPString removes all leading and trailing white space and replace quotation marks with &#34;
func HTTPString(input string) (output String) {
	if len(input) != 0 {
		return String{strings.Replace(strings.TrimSpace(input), "\"", "&#34;", -1), true}
	}
	return
}

// MultipleEqual checks all the bool parameters and returns a result
func MultipleEqual(bools ...bool) (bool, error) {

	if len(bools) > 255 {
		return false, errors.New("the number of input values exceeds 255")
	}

	var equal bool = true
	var length = byte(len(bools))

	if length < 2 {
		return false, errors.New("the number of input values is less then 2")
	}

	var i byte = 1
	for ; equal && i < length; i++ {
		equal = bools[i] == bools[i-1]
	}

	return equal, nil
}

// CompareTwoStrings compares two string
func CompareTwoStrings(string1, string2 string) bool {
	return string1 == string2
}

// Days returns number of days in a month
func Days(month time.Time) int {
	month = month.AddDate(0, 1, 0)
	timeToGetLastDay := Date(month.Year(), 0, month.Month())
	return timeToGetLastDay.Day()
}

// Date is a shorter version of the time.Date() function
func Date(year, day int, month time.Month) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

// IsEmailValid checks if the input email is valid.
func IsEmailValid(email string) bool {
	if len(email) < 3 && len(email) > 254 {
		return false
	}
	return regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$").MatchString(email)
}
