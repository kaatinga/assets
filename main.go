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

// Uint16 is an extended version of uint16 type
type Uint16 struct {
	Parameter uint16
	Ok        bool
}

func (Uint16 *Uint16) IsOk() bool {
	return Uint16.Ok
}

func (Uint16 *Uint16) Unpack() (uint16, bool) {
	return Uint16.Parameter, Uint16.Ok
}

// CheckUint16 checks and converts input string to uint16 type
func CheckUint16(inputString string) (output Uint16) {
	var err error  // to store error result
	var tmpInt int // a temporary int value

	tmpInt, err = strconv.Atoi(inputString)
	if err == nil {
		if tmpInt >= 0 && tmpInt < 65536 {
			output.Parameter = uint16(tmpInt)
			output.Ok = true
			return output
		}
	}
	return output
}

// SetUint16 checks and sets input string to uint16 type
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
func CheckName(name string) (ok bool) {
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

// A generic type
type Gen struct {
	Parameter interface{}
	Ok        bool
}

// Uint16 checks and sets uint16 value from generic to a pointer to uint16
func (gen Gen) Uint16(parameter *uint16) bool {

	switch gen.Parameter.(type) {
	case uint16:
	default:
		return false
	}

	if gen.Ok {
		*parameter = gen.Parameter.(uint16)
		return true
	}
	return false
}

// Byte checks and sets uint16 value from generic to a pointer to byte
func (gen Gen) Byte(parameter *byte) bool {

	switch gen.Parameter.(type) {
	case byte:
	default:
		return false
	}

	if gen.Ok {
		*parameter = gen.Parameter.(byte)
		return true
	}
	return false
}

func CheckGen(input ...interface{}) (gen Gen) {

	switch input[1].(type) {
	case bool:
		if !input[1].(bool) {
			return
		}
	default:
		gen.Ok = false
		return
	}

	switch input[0].(type) {
	case uint16, byte:
		gen.Parameter = input[0]
		gen.Ok = true
	default:
		gen.Ok = false
		return
	}

	return
}
