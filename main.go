package assets

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"
)

const (
	maxUint8  uint64 = 255
	maxUint16 uint64 = 65535
	maxUint32 uint64 = 4294967295
)

// SuperBytesToUint32 checks and converts input string to uint32 type
func SuperBytesToUint32(input []byte) (output uint32, ok bool) {

	var output64 uint64

	log.Println("строчное представление input", string(input))
	log.Println("input", input)

	if input[0] < 48 || input[0] > 57 {
		return
	}

MainLoop:
	for key, value := range input {
		log.Println("=== байт", value)

		switch key {
		case 0:
			log.Println("преобразуем первый байт")
			output64 = uint64(value) - 48
			log.Println("промежуточный итог", output64)
			continue MainLoop
		default:
			if value < 48 || value > 57 {
				log.Println("это неправильное значение", value)
				ok = true
				break MainLoop
			}
		}

		log.Println("умножаем на 10 и прибавляем")
		output64 = output64*10 + uint64(value) - 48

		if output64 >= maxUint32 {
			log.Println("цифра слишком большая!")
			return 0, false
		}

		log.Println("промежуточный итог", output64)
	}

	output = uint32(output64)
	ok = true
	log.Println("окончательный итог", output)
	return
}

// StByte checks and converts input string to Byte type
func StByte(inputString string) (byte, bool) {
	var (
		tmpUint64 uint64 // a temporary int value
		ok        bool
	)

	tmpUint64, ok = StUint64(inputString)
	if ok {
		if tmpUint64 >= 0 && tmpUint64 <= maxUint8 {
			return byte(tmpUint64), true
		}
	}
	return 0, false
}

// StUint16 checks and converts input string to uint16 type
func StUint16(inputString string) (uint16, bool) {
	var (
		tmpUint64 uint64 // a temporary int value
		ok        bool
	)

	tmpUint64, ok = StUint64(inputString)
	if ok {
		if tmpUint64 >= 0 && tmpUint64 <= maxUint16 {
			return uint16(tmpUint64), true
		}
	}

	return 0, false
}

// StUint32 checks and converts input string to uint16 type
func StUint32(inputString string) (uint32, bool) {
	var (
		tmpUint64 uint64 // a temporary int value
		ok        bool
	)

	tmpUint64, ok = StUint64(inputString)
	if ok {
		if tmpUint64 >= 0 && tmpUint64 <= maxUint32 {
			return uint32(tmpUint64), true
		}
	}

	return 0, false
}

// StUint64 checks and converts input string to uint64 type
func StUint64(inputString string) (output uint64, ok bool) {
	var err error // to store error result

	output, err = strconv.ParseUint(inputString, 10, 64)
	return output, err == nil
}

// Uint16 is an extended version of uint16 type
type Uint16 struct {
	Parameter uint16
	Ok        bool
}

// CheckUint16 checks and converts input string to Uint16 struct
func CheckUint16(inputString string) (output Uint16) {
	output.Parameter, output.Ok = StUint16(inputString)
	return
}

// SetByte checks and sets input string to a given pointer to byte type
func SetByte(inputByte *byte, inputString string) (ok bool) {

	*inputByte, ok = StByte(inputString)
	return
}

// SetUint16 checks and sets input string to a given pointer to uint16 type
func SetUint16(inputUint16 *uint16, inputString string) (ok bool) {

	*inputUint16, ok = StUint16(inputString)
	return
}

// SetUint32 checks and sets input string to a given pointer to uint16 type
func SetUint32(inputUint32 *uint32, inputString string) (ok bool) {

	*inputUint32, ok = StUint32(inputString)
	return
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
