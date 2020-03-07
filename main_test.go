package assets

import (
	"log"
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
		{"incorrect 7", args{7}, 6, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStr, err := GenPassword(tt.args.length)
			log.Println(gotStr)
			if (err != nil) != tt.wantErr {
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
