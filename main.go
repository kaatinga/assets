package assets

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"regexp"
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

func init() {
	// in order to issue really random password
	rand.Seed(time.Now().UnixNano())
}

// StByte checks and converts input string to Byte type
func StByte(inputString string) (byte, bool) {
	var (
		tmpUint64 uint64 // a temporary int value
		ok        bool
	)

	tmpUint64, ok = StUint64(inputString)
	if ok && tmpUint64 <= maxUint8 {
		return byte(tmpUint64), true
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
	if ok && tmpUint64 <= maxUint16 {
		return uint16(tmpUint64), true
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
	if ok && tmpUint64 <= maxUint32 {
		return uint32(tmpUint64), true
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

// SetUint16 checks and sets input string to a given pointer to uint16 type.
func SetUint16(inputUint16 *uint16, inputString string) (ok bool) {

	*inputUint16, ok = StUint16(inputString)
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

// a random number generator that returns a byte
func GetRandomByte(max byte) byte {

	switch max {
	case 0:
		return 0
	default:
		return byte(rand.Int31n(int32(max)))
	}
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

// SafeQM escapes quotation marks adding '\' before them
func SafeQM(str string) (newString string) {
	newString = strings.Replace(str, `"`, `\"`, -1)
	return
}

// RemoveSafeQM removes symbols '\' before quotation marks
func RemoveSafeQM(str string) (newString string) {
	newString = strings.Replace(str, `\"`, `"`, -1)
	return
}

// CheckRussianCompanyName check if an only allowed set of symbols is in the company name
func CheckRussianCompanyName(company string) (ok bool) {

	// Russian company can have digits and russian symbols, as well as some symbols below
	var symbolRange = []*unicode.RangeTable{
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
	switch input {
	case "":
	default:
		output.Ok = true
		output.Parameter = strings.Replace(strings.TrimSpace(input), "\"", "&#34;", -1)
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

// Checks if the input email is valid
func IsEmailValid(email string) bool {
	if len(email) < 3 && len(email) > 254 {
		return false
	}
	return regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$").MatchString(email)
}

var uint16sizes = []uint16{
	1: 10,
	2: 100,
	3: 1000,
	4: 10000,
}

const (
	ten         uint16 = 10
	hundred     uint16 = 100
	thousand    uint16 = 1000
	tenThousand uint16 = 10000
)

// Uint162Bytes converts an uint16 number to string.
func Uint162Bytes(num uint16) []byte {

	convertedNumber, i := getSliceUint16(num)

	for {
		convertedNumber[i] = byte(num%10) + 0x30
		num = num / 10
		if i == 0 {
			return convertedNumber
		}
		i--
	}
}

// Uint162String converts an uint16 number to string.
func Uint162String(num uint16) string {

	convertedNumber, i := getSliceUint16(num)

	for {
		convertedNumber[i] = byte(num%10) + 0x30
		num = num / 10
		if i == 0 {
			return string(convertedNumber)
		}
		i--
	}
}

func getSliceUint16(num uint16) ([]byte, int) {
	if num < ten {
		return make([]byte, 1), 0
	}

	if num < hundred {
		return make([]byte, 2), 1
	}

	if num < thousand {
		return make([]byte, 3), 2
	}

	if num < tenThousand {
		return make([]byte, 4), 3
	}

	return make([]byte, 5), 4
}

//func getSize(num uint16, size int) int {
//	if size == 5 {
//		return size
//	}
//	if num < uint16sizes[size] {
//		return size
//	}
//	return getSize(num, size+1)
//}
//
//func convertNumber(convertedNumber []byte, i int, num uint16) {
//	convertedNumber[i] = byte(num%10) + 0x30
//	num = num / 10
//	if num == 0 {
//		return
//	}
//	convertNumber(convertedNumber, i-1, num)
//}

// Byte2Bytes converts a byte number to []byte.
func Byte2Bytes(num byte) []byte {

	convertedNumber, i := getSliceByte(num)

	for {
		convertedNumber[i] = num%10 + 0x30
		num = num / 10
		if i == 0 {
			return convertedNumber
		}
		i--
	}
}

// Byte2String converts a byte number to string.
func Byte2String(num byte) string {

	convertedNumber, i := getSliceByte(num)

	for {
		convertedNumber[i] = num%10 + 0x30
		num = num / 10
		if i == 0 {
			return string(convertedNumber)
		}
		i--
	}
}

const (
	byteTen     byte = 10
	byteHundred byte = 100
)

func getSliceByte(num byte) ([]byte, int) {
	if num < byteTen {
		return make([]byte, 1), 0
	}

	if num < byteHundred {
		return make([]byte, 2), 1
	}

	return make([]byte, 3), 2
}
