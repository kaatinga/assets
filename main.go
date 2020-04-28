package assets

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// StUint16 checks and converts input string to uint16 type
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

// CheckUint16 checks and converts input string to uint16 type. Return Uint16 struct
func CheckUint16(inputString string) (output Uint16) {
	output.Parameter, output.Ok = StUint16(inputString)
	return
}

// StUint32 converts input string to uint16 type
func StUint32(inputString string) (uint32, bool) {
	var err error  // to store error result
	var tmpInt int // a temporary int value

	tmpInt, err = strconv.Atoi(inputString)
	if err == nil {
		if tmpInt >= 0 && tmpInt < 4294967296 {
			return uint32(tmpInt), true
		}
	}
	return 0, false
}

// Uint16 is an extended version of uint16 type
type Uint16 struct {
	Parameter uint16
	Ok        bool
}

// String is an extended version of uint16 type
type String struct {
	Parameter string
	Ok        bool
}

// SetUint16 checks and sets input string to a given uint16 type
func SetUint16(inputUint16 *uint16, inputString string) bool {
	var err error  // to store error result
	var tmpInt int // a temporary int value

	tmpInt, err = strconv.Atoi(inputString)
	if err == nil {
		if tmpInt >= 0 && tmpInt < 65536 {
			*inputUint16 = uint16(tmpInt)
			return true
		}
	}
	return false
}

// SetUint32 checks and sets input string to uint16 type
func SetUint32(inputUint32 *uint32, inputString string) bool {
	var err error  // to store error result
	var tmpInt int // a temporary int value

	tmpInt, err = strconv.Atoi(inputString)
	if err == nil {
		if tmpInt >= 0 && tmpInt < 4294967296 {
			*inputUint32 = uint32(tmpInt)
			return true
		}
	}
	return false
}

// SetStringByPointer checks and sets input Strings parameter to a string through a pointer
func (input *String) SetStringByPointer(output *string) bool {
	if (*input).Parameter != "" {
		*output = (*input).Parameter
		return true
	}
	return false
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

// SetByte checks and sets input string to byte type
func SetByte(inputByte *byte, inputString string) bool {
	var err error  // to store error result
	var tmpInt int // a temporary int value

	tmpInt, err = strconv.Atoi(inputString)
	if err == nil {
		if tmpInt >= 0 && tmpInt < 256 {
			*inputByte = byte(tmpInt)
			return true
		}
	}
	return false
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
		if err != nil {
			return "", err
		}
	}
	return builder.String(), nil
}

// SaveFile saves a file
func SaveFile(data interface{}, path string) (err error) {
	var dataToWrite []byte // the variable to store serialized JSON data
	dataToWrite, err = json.Marshal(data)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, dataToWrite, 0660)
	if err != nil {
		return err
	}

	return nil
}

// SafeQM escapes quatation marks adding '\' before them
func SafeQM(str string) (newString string) {
	newString = strings.Replace(str, `"`, `\"`, -1)
	return
}

// RemoveSafeQM removes symbols '\' before quatation marks
func RemoveSafeQM(str string) (newString string) {
	newString = strings.Replace(str, `\"`, `"`, -1)
	return
}

// CheckRussianCompanyName check if an only allowed set of symbols is in the company name
func CheckRussianCompanyName(company string) (ok bool) {
	// Russian company can have digits and russian symbols, as well as soon symbols below
	var symbolRange []*unicode.RangeTable = []*unicode.RangeTable{
		unicode.Cyrillic,
		unicode.Digit,
	}

	company = RemoveCharacters(company, "& \"+-»«") // to remove a set of allowed symbols
	companyRune := []rune(company)

	for _, value := range companyRune {
		if !unicode.IsOneOf(symbolRange, value) {
			return false
		}
	}

	return true
}

// CheckName check if an only allowed set of symbols is in the string
func CheckName(name string) bool {
	name = RemoveCharacters(name, " ") // to remove space
	nameRune := []rune(name)

	for _, value := range nameRune {
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
	if input == "" {
		output.Ok = false
		return output
	}
	output.Ok = true
	output.Parameter = strings.TrimSpace(input)
	output.Parameter = strings.Replace(output.Parameter, "\"", "&#34;", -1)
	return
}

// MultipleEqual checks all the bool parameters and returns a result
func MultipleEqual(bools ...bool) byte {

	var previous byte
	var previousBool bool
	for _, value := range bools {
		if previous == 0 {
			previousBool = value
			if value == true {
				previous = 2
			}

			previous = 1
			continue
		}

		if previousBool != value {
			return 3 // 3 means the values are not equal
		}
	}

	if previousBool == false {
		return 1 // 1 = false
	}

	return 2 // 2 = true
}

// CompareTwoStrings compares two string
func CompareTwoStrings(string1, string2 string) bool {
	return string1 == string2
}
