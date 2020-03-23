package assets

import (
	"reflect"
	"testing"
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStr, err := GenPassword(tt.args.length)
			ok := (err != nil)
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StBool(tt.args.inputString); got != tt.want {
				t.Errorf("StBool() = %v, want %v", got, tt.want)
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
		{"correct string", args{"9"}, 9, true},
		{"zero", args{"0"}, 0, true},
		{"negatrive string", args{"-9"}, 0, false},
	}
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
		{"correct string", args{"9"}, 9, true},
		{"negatrive string", args{"-9"}, 0, false},
	}
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

func TestSaveFile(t *testing.T) {
	type args struct {
		data interface{}
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SaveFile(tt.args.data, tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("SaveFile() error = %v, wantErr %v", err, tt.wantErr)
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
		{`englishstring`, args{"ООО «Company»"}, false},
	}
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := HTTPString(tt.args.input); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("HTTPString() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestString_IsOk(t *testing.T) {
	correctString := String{"&#34;abc&#34;", true}
	incorrectString := String{"", false}
	tests := []struct {
		name   string
		String *String
		want   bool
	}{
		{`correct`, &correctString, true},
		{`incorrect`, &incorrectString, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.String.IsOk(); got != tt.want {
				t.Errorf("String.IsOk() = %v, want %v", got, tt.want)
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.input.SetStringByPointer(tt.args.output); got != tt.want {
				t.Errorf("String.SetStringByPointer() = %v, want %v", got, tt.want)
			}
		})
	}
}
