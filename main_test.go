package assets

import (
	"reflect"
	"testing"
	"time"
)

func TestGenPassword(t *testing.T) {
	type args struct {
		length byte
	}
	tests := []struct {
		name      string
		args      args
		strLength int
		wantErr   bool
	}{
		{"Zero value", args{0}, 7, false},
		{"6", args{6}, 6, false},
		{"16", args{16}, 16, false},
	}

	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStr, err := GenPassword(tt.args.length)
			ok := err != nil
			if ok != tt.wantErr {
				t.Errorf("GenPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(gotStr) != tt.strLength {
				t.Errorf("GenPassword() gotStr = %v, want %v", gotStr, tt.strLength)
			}
		})
	}
}

func TestStBool(t *testing.T) {
	type args struct {
		inputString string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"true 1", args{"on"}, true},
		{"true 2", args{"true"}, true},
		{"false 1", args{""}, false},
		{"false 2", args{"да!"}, false},
	}

	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StBool(tt.args.inputString); got != tt.want {
				t.Errorf("StBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStUint64(t *testing.T) {
	type args struct {
		inputString string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput uint64
		wantOk     bool
	}{
		{"correct string 1", args{"99999"}, 99999, true},
		{"correct string 2", args{"65535"}, 65535, true},
		{"correct string 3", args{"9"}, 9, true},
		{"negative string", args{"-9"}, 0, false},
	}

	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOutput, gotOk := StUint64(tt.args.inputString)
			if gotOutput != tt.wantOutput {
				t.Errorf("StUint64() gotOutput = %v, want %v", gotOutput, tt.wantOutput)
			}
			if gotOk != tt.wantOk {
				t.Errorf("StUint64() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestStByte(t *testing.T) {
	type args struct {
		inputString string
	}
	tests := []struct {
		name  string
		args  args
		want  byte
		want1 bool
	}{
		{"incorrect string", args{"9999"}, 0, false},
		{"correct string 1", args{"9"}, 9, true},
		{"correct string 2", args{"255"}, 255, true},
		{"zero", args{"0"}, 0, true},
		{"negative string", args{"-9"}, 0, false},
	}

	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := StByte(tt.args.inputString)
			if got != tt.want {
				t.Errorf("StByte() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("StByte() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestStUint16(t *testing.T) {
	type args struct {
		inputString string
	}
	tests := []struct {
		name  string
		args  args
		want  uint16
		want1 bool
	}{
		{"incorrect string", args{"99999"}, 0, false},
		{"correct string 1", args{"65535"}, 65535, true},
		{"correct string 2", args{"9"}, 9, true},
		{"negative string", args{"-9"}, 0, false},
	}

	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := StUint16(tt.args.inputString)
			if got != tt.want {
				t.Errorf("StUint16() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("StUint16() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSafeQM(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name          string
		args          args
		wantNewString string
	}{
		{`correct string`, args{`ООО "Ромашка"`}, `ООО \"Ромашка\"`},
	}

	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewString := SafeQM(tt.args.str); gotNewString != tt.wantNewString {
				t.Errorf("SafeQM() = %v, want %v", gotNewString, tt.wantNewString)
			}
		})
	}
}

func TestRemoveSafeQM(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name          string
		args          args
		wantNewString string
	}{
		{`correct string`, args{`ООО \"Ромашка\"`}, `ООО "Ромашка"`},
	}

	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewString := RemoveSafeQM(tt.args.str); gotNewString != tt.wantNewString {
				t.Errorf("RemoveSafeQM() = %v, want %v", gotNewString, tt.wantNewString)
			}
		})
	}
}

func TestCheckRussianCompanyName(t *testing.T) {
	type args struct {
		company string
	}
	tests := []struct {
		name   string
		args   args
		wantOk bool
	}{
		{`string1`, args{"ООО «аб_в»"}, false},
		{`string+digits`, args{"ООО «1а2б_3в»"}, false},
		{`string2`, args{"ООО «Про&+ба»"}, true},
		{`string+digits2`, args{"ООО «а-б1-в»"}, true},
		{`string3`, args{"ООО \"а&бв\""}, true},
		{`english string`, args{"ООО «Company»"}, false},
	}

	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOk := CheckRussianCompanyName(tt.args.company); gotOk != tt.wantOk {
				t.Errorf("CheckCompanyName() = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func Test_removeCharacters(t *testing.T) {
	type args struct {
		input      string
		characters string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{`string1`, args{"ООО «аб_в»", "&\"+-_»«"}, "ООО абв"},
	}

	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveCharacters(tt.args.input, tt.args.characters); got != tt.want {
				t.Errorf("removeCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetUint16(t *testing.T) {
	var testUint16 uint16
	type args struct {
		inputUint16 *uint16
		inputString string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{`correct`, args{&testUint16, "16"}, true},
		{`incorrect`, args{&testUint16, "-16"}, false},
	}

	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetUint16(tt.args.inputUint16, tt.args.inputString); got != tt.want {
				t.Errorf("SetUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPString(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput String
	}{
		{`correct`, args{` "abc" `}, String{"&#34;abc&#34;", true}},
		{`incorrect`, args{""}, String{"", false}},
	}

	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := HTTPString(tt.args.input); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("HTTPString() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestString_SetStringByPointer(t *testing.T) {
	correctString := String{"&#34;abc&#34;", true}
	var tempString string
	type args struct {
		output *string
	}
	tests := []struct {
		name  string
		input *String
		args  args
		want  bool
	}{
		{`correct`, &correctString, args{&tempString}, true},
	}

	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.SetStringByPointer(tt.args.output); got != tt.want {
				t.Errorf("String.SetStringByPointer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultipleEqual(t *testing.T) {
	type args struct {
		bools []bool
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{`all true`, args{[]bool{true, true, true, true}}, true, false},
		{`all false`, args{[]bool{false, false, false, false}}, true, false},
		{`true and false`, args{[]bool{true, true, true, false}}, false, false},
		{`too short`, args{[]bool{true}}, false, true},
	}

	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := MultipleEqual(tt.args.bools...)
			if tt.wantErr != (err != nil) {
				t.Errorf("MultipleEqual() returned error, but did not have to")
			}

			if got != tt.want {
				t.Errorf("MultipleEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompareTwoStrings(t *testing.T) {
	type args struct {
		string1 string
		string2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{`true`, args{"a", "a"}, true},
		{`false`, args{"a", "b"}, false},
	}

	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareTwoStrings(tt.args.string1, tt.args.string2); got != tt.want {
				t.Errorf("CompareTwoStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckUint16(t *testing.T) {
	type args struct {
		inputString string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput Uint16
	}{
		{`false 1`, args{"a"}, Uint16{0, false}},
		{`true 1`, args{"0"}, Uint16{0, true}},
		{`true 2`, args{"55"}, Uint16{55, true}},
		{`true 2`, args{"65535"}, Uint16{65535, true}},
		{`false 2`, args{"65536"}, Uint16{0, false}},
	}

	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := CheckUint16(tt.args.inputString); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("CheckUint16() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestStUint32(t *testing.T) {
	type args struct {
		inputString string
	}
	tests := []struct {
		name  string
		args  args
		want  uint32
		want1 bool
	}{
		{"incorrect string", args{"4294967297"}, 0, false},
		{"correct string 1 ", args{"429496725"}, 429496725, true},
		{"correct string 2", args{"9"}, 9, true},
		{"negative string", args{"-9"}, 0, false},
		{"zero", args{"0"}, 0, true},
	}

	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := StUint32(tt.args.inputString)
			if got != tt.want {
				t.Errorf("StUint32() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("StUint32() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCheckName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"ok", args{"Русское слово"}, true},
		{"english", args{"Nerusskoe слово"}, false},
		{"123", args{"123"}, false},
	}

	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckName(tt.args.name); got != tt.want {
				t.Errorf("CheckName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSuperBytesToUint32(t *testing.T) {

	tests := []struct {
		name       string
		args       []byte
		wantOutput uint32
		wantErr     bool
	}{
		{"correct string 42", []byte("42"), 42, true},
		{"correct string 12", []byte("12"), 12, true},
		//{"correct input []byte{49, 50, 13, 10}", []byte{49, 50, 13, 10}, 0, true},
		{"correct input []byte{48})", []byte{48}, 0, true},
		{"incorrect input []byte{13, 10}", []byte{13, 10}, 0, false},
		{"incorrect input []byte{0}", []byte{0}, 0, false},
		{"incorrect string 4294967296", []byte("4294967296"), 0, false},
		{"incorrect string -1", []byte("-1"), 0, false},
	}

	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOutput, err := SuperBytesToUint32(tt.args)
			if gotOutput != tt.wantOutput {
				t.Errorf("SuperBytesToUint32() gotOutput = %v, want %v", gotOutput, tt.wantOutput)
			}
			if (err != nil) == tt.wantErr {
				t.Errorf("SuperBytesToUint32() got error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getRandomByte(t *testing.T) {

	tests := []struct {
		name string
		max  byte
	}{
		{"ok1", 255},
		{"ok2", 100},
		{"ok3", 100},
		{"ok4", 100},
		{"ok5", 100},
		{"ok6", 100},
		{"ok7", 100},
		{"ok8", 1},
		{"ok9", 0},
	}

	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetRandomByte(tt.max)

			if got > tt.max {
				t.Errorf("getRandomByte() got = %v, wants a number not bigger than %v", got, tt.max)
			}
		})
	}
}

func TestDays(t *testing.T) {

	goodTime1, _ := time.Parse(time.RFC1123, "Wed, 02 Dec 2020 00:00:00 UTC")
	goodTime2, _ := time.Parse(time.RFC1123, "Wed, 01 Feb 2000 00:00:00 UTC")

	tests := []struct {
		name  string
		month time.Time
		want  int
	}{
		{"ok1", goodTime1, 31},
		{"ok2", goodTime2, 29},
	}

	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Days(tt.month); got != tt.want {
				t.Errorf("Days() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsEmailValid(t *testing.T) {

	tests := []struct {
		name  string
		email string
		want  bool
	}{
		{"ok", "test@golangcode.com", true},
		{"!ok1", "test", false},
		{"!ok2", "123@1???--23", false},
	}

	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmailValid(tt.email); got != tt.want {
				t.Errorf("IsEmailValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint162String(t *testing.T) {

	tests := []struct {
		result string
		num    uint16
	}{
		{"199", 199},
		{"1999", 1999},
		{"222", 222},
		{"1", 1},
		{"0", 0},
		{"55555", 55555},
		{"12345", 12345},
		{"10000", 10000},
		{"9999", 9999},
	}

	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.result, func(t *testing.T) {
			got := Uint162String(tt.num)
			if got != tt.result {
				t.Errorf("Uint162String() = %v, want %v", got, tt.result)
			}
		})
	}
}

func TestByte2String(t *testing.T) {

	tests := []struct {
		result string
		num    byte
	}{
		{"199", 199},
		{"99", 99},
		{"100", 100},
		{"255", 255},
		{"222", 222},
		{"0", 0},
		{"9", 9},
		{"10", 10},
		{"1", 1},
	}

	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.result, func(t *testing.T) {
			if got := Byte2String(tt.num); !reflect.DeepEqual(got, tt.result) {
				t.Errorf("Byte2String() = %v, want %v", got, tt.result)
			}
		})
	}
}
