package assets

import (
	"strconv"
	"testing"
)

func TestString2Uint32(t *testing.T) {

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
		{"incorrect string 25549672951", "25549672951", 0, ErrNumberExceedMaxUint32Value, ErrNotByte},
		{"incorrect string 429496729612", "429496729612", 0, ErrNumberExceedMaxUint32Value, ErrNotByte},
		{"incorrect string -1", "-1", 0, ErrNotUint32, ErrNotByte},
	}

	gotOutputByte, err := String2Byte("257")
	if err != ErrNumberExceedMaxByteValue {
		t.Errorf("String2Byte() error\nhave '%v'\nwant '%v'", err, ErrNumberExceedMaxByteValue)
		t.Log("got number:", gotOutputByte, "sample: 257")
		return
	}
	if gotOutputByte != 0 {
		t.Errorf("String2Byte() gotOutput = %v, want %v", gotOutputByte, 0)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotOutput uint32
			gotOutput, err = String2Uint32(tt.input)
			if err != tt.wantErrUint32 {
				t.Errorf("String2Uint32() error = %v, want %v", err, tt.wantErrUint32)
			}
			if gotOutput != tt.wantOutput {
				t.Errorf("String2Uint32() gotOutput = %v, want %v", gotOutput, tt.wantOutput)
			}

			gotOutput, err = Bytes2Uint32([]byte(tt.input))
			if err != tt.wantErrUint32 {
				t.Errorf("Bytes2Uint32() error = %v, want %v", err, tt.wantErrUint32)
			}
			if gotOutput != tt.wantOutput {
				t.Errorf("Bytes2Uint32() gotOutput = %v, want %v", gotOutput, tt.wantOutput)
			}

			if tt.wantOutput < 255 {
				gotOutputByte, err = String2Byte(tt.input)
				if err != tt.wantErrByte {
					t.Errorf("String2Byte() error\nhave '%v'\nwant '%v'", err, tt.wantErrByte)
					t.Log(gotOutputByte)
					return
				}
				if gotOutputByte != byte(tt.wantOutput) {
					t.Errorf("String2Byte() gotOutput = %v, want %v", gotOutputByte, tt.wantOutput)
					t.FailNow()
				} else {
					t.Log("byte is ok!")
				}
			}
		})
	}
}

func TestString2Byte(t *testing.T) {

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

		//fmt.Println("ok!", i)
	}
}
