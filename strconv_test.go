package assets

import (
	"strconv"
	"testing"
)

func Test2Uint32(t *testing.T) {

	tests := []struct {
		name          string
		input         string
		wantOutput    uint32
		wantErrUint32 error
		wantErrByte   error
	}{
		{"correct string 42", "42", 42, nil, nil},
		{"correct string 12", "12", 12, nil, nil},
		//{"correct input 012", "012", 12, nil, nil},
		{"correct input 0", "0", 0, nil, nil},
		{"correct string 4294967295", "4294967295", 4294967295, nil, nil},
		{"incorrect input []byte{13, 10}", string([]byte{13, 10}), 0, ErrNotUint32, ErrNotByte},
		{"incorrect input []byte{0}", string([]byte{0}), 0, ErrNotUint32, ErrNotByte},
		{"incorrect string 4294967296", "4294967296", 0, ErrNumberExceedMaxUint32Value, ErrNotByte},
		{"incorrect string 25549672951", "25549672951", 0, ErrNotUint32, ErrNotByte},
		{"incorrect string 429496729612", "429496729612", 0, ErrNotUint32, ErrNotByte},
		{"incorrect string -1", "-1", 0, ErrNotUint32, ErrNotByte},
		{"incorrect empty string", "", 0, ErrNotUint32, ErrNotByte},
	}

	var err error

	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotOutput uint32
			gotOutput, err = String2Uint32(tt.input)
			if err != tt.wantErrUint32 { //nolint:errorlint
				t.Errorf("String2Uint32() error = %v, want %v", err, tt.wantErrUint32)
			}
			if gotOutput != tt.wantOutput { //nolint:errorlint
				t.Errorf("String2Uint32() gotOutput = %v, want %v", gotOutput, tt.wantOutput)
			}

			gotOutput, err = Bytes2Uint32([]byte(tt.input))
			if err != tt.wantErrUint32 { //nolint:errorlint
				t.Errorf("Bytes2Uint32() error = %v, want %v", err, tt.wantErrUint32)
			}
			if gotOutput != tt.wantOutput { //nolint:errorlint
				t.Errorf("Bytes2Uint32() gotOutput = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func Test2Uint32_2(t *testing.T) {

	var stringValue string
	var value int
	var value3 uint32
	var err error
	for i := uint64(0); i < 10000000000; i = i + 10001 {
		stringValue = strconv.FormatUint(i, 10)
		value, _ = strconv.Atoi(stringValue)

		value3, err = String2Uint32(stringValue)
		if i < 4294967296 {
			if err != nil {
				t.Errorf("String2Uint32() error = %v, want %v", err, nil)
			}

			if uint32(value) != value3 {
				t.Errorf("String2Uint32() value = %v, want %v", value3, value)
			}
		} else {
			if err == nil {
				t.Error("String2Uint32() must have an error")
			}

			if 0 != value3 {
				t.Errorf("String2Uint32() value = %v, want 0, i = %v", value3, i)
			}
		}

		value3, err = Bytes2Uint32([]byte(stringValue))
		if i < 4294967296 {
			if err != nil {
				t.Errorf("Bytes2Uint32() error = %v, want %v", err, nil)
			}

			if uint32(value) != value3 {
				t.Errorf("Bytes2Uint32() value = %v, want %v", value3, value)
			}
		} else {
			if err == nil {
				t.Error("Bytes2Uint32() must have an error")
			}

			if 0 != value3 {
				t.Errorf("Bytes2Uint32() value = %v, want %v, i %v", value3, 0, i)
			}
		}
	}
}

func Test2Byte(t *testing.T) {

	var stringValue string
	var value int
	var value3 byte
	var err error
	for i := 0; i < 1000; i++ {
		stringValue = strconv.Itoa(i)
		value, _ = strconv.Atoi(stringValue)

		value3, err = String2Byte(stringValue)
		if i < 256 {
			if err != nil {
				t.Errorf("String2Byte() error = %v, want %v", err, nil)
			}

			if byte(value) != value3 {
				t.Errorf("String2Byte() value = %v, want %v", value3, value)
			}
		} else {
			if err == nil {
				t.Error("String2Byte() must have an error")
			}

			if 0 != value3 {
				t.Errorf("String2Byte() value = %v, want %v, i %v", value3, 0, i)
			}
		}

		value3, err = Bytes2Byte([]byte(stringValue))
		if i < 256 {
			if err != nil {
				t.Errorf("Bytes2Byte() error = %v, want %v", err, nil)
			}

			if byte(value) != value3 {
				t.Errorf("Bytes2Byte() value = %v, want %v", value3, value)
			}
		} else {
			if err == nil {
				t.Error("Bytes2Byte() must have an error")
			}

			if 0 != value3 {
				t.Errorf("Bytes2Byte() value = %v, want %v, i %v", value3, 0, i)
			}
		}
	}
}

func TestBytes2Uint16(t *testing.T) {

	tests := []struct {
		input   []byte
		want    uint16
		wantErr bool
	}{
		{[]byte("255"), 255, false},
		{[]byte("25"), 25, false},
		{[]byte("0"), 0, false},
		{[]byte("16000"), 16000, false},
	}
	for _, tt := range tests {
		t.Run(string(tt.input), func(t *testing.T) {
			got, err := Bytes2Uint16(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Bytes2Uint16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Bytes2Uint16() got = %v, want %v", got, tt.want)
			}
		})
	}
}
