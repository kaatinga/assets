package assets

import (
	"reflect"
	"testing"
	"time"
)

func TestStBool(t *testing.T) {

	tests := []struct {
		inputString string
		want        bool
	}{
		{"TruE", true},
		{"true", true},
		{"True", true},
		{"TRUE", true},
		{"", false},
		{"труе", false},
		{"да!", false},
	}

	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.inputString, func(t *testing.T) {
			if got := String2Bool(tt.inputString); got != tt.want {
				t.Errorf("String2Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeQM(t *testing.T) {

	tests := []struct {
		name          string
		str           string
		wantNewString string
	}{
		{`correct string`, `ООО "Ромашка"`, `ООО \"Ромашка\"`},
	}

	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewString := SafeQM(tt.str); gotNewString != tt.wantNewString {
				t.Errorf("SafeQM() = %v, want %v", gotNewString, tt.wantNewString)
			}
		})
	}
}

func TestRemoveSafeQM(t *testing.T) {

	tests := []struct {
		name          string
		str           string
		wantNewString string
	}{
		{`correct string`, `ООО \"Ромашка\"`, `ООО "Ромашка"`},
	}

	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewString := RemoveSafeQM(tt.str); gotNewString != tt.wantNewString {
				t.Errorf("RemoveSafeQM() = %v, want %v", gotNewString, tt.wantNewString)
			}
		})
	}
}

func TestCheckRussianCompanyName(t *testing.T) {
	tests := []struct {
		name    string
		company string
		wantOk  bool
	}{
		{`string1`, "ООО «аб_в»", false},
		{`string+digits`, "ООО «1а2б_3в»", false},
		{`string2`, "ООО «Про&+ба»", true},
		{`string+digits2`, "ООО «а-б1-в»", true},
		{`string3`, "ООО \"а&бв\"", true},
		{`english string`, "ООО «Company»", false},
		{`english string`, "ООО «Ромашка №1»", true},
	}

	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOk := CheckRussianCompanyName(tt.company); gotOk != tt.wantOk {
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
