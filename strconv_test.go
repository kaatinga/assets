package assets

import "testing"

func TestString2Uint32(t *testing.T) {

	tests := []struct {
		name       string
		input      string
		wantOutput uint32
		wantErr    error
	}{
		{"correct string 42", "42", 42, nil},
		{"correct string 12", "12", 12, nil},
		{"correct input 012", "012", 12, nil},
		{"correct input 0)", "0", 0, nil},
		{"incorrect input []byte{13, 10}", string([]byte{13, 10}), 0, ErrNotUint32},
		{"incorrect input []byte{0}", string([]byte{0}), 0, ErrNotUint32},
		{"incorrect string 4294967296", "4294967296", 0, ErrNumberExceedMaxUint32Value},
		{"incorrect string 429496729612", "429496729612", 0, ErrNumberExceedMaxUint32Value},
		{"incorrect string -1", "-1", 0, ErrNotUint32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOutput, err := String2Uint32(tt.input)
			if err != tt.wantErr {
				t.Errorf("String2Uint32() error = %v, wantErr %v", err, tt.wantErr)
				t.Log(gotOutput)
				return
			}
			if gotOutput != tt.wantOutput {
				t.Errorf("String2Uint32() gotOutput = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
