package assets

import (
	"testing"
)

func TestGenPassword(t *testing.T) {

	tests := []struct {
		name      string
		length    byte
		strLength int
		wantErr   bool
	}{
		{"Zero value", 0, 7, false},
		{"6", 6, 6, false},
		{"16", 16, 16, false},
	}

	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStr := GenPassword(tt.length)
			t.Log(gotStr)
			if len(gotStr) != tt.strLength {
				t.Errorf("GenPassword() gotStr = %v, want %v", gotStr, tt.strLength)
			}
		})
	}
}

func TestGenPasswordAsBytes(t *testing.T) {

	tests := []struct {
		name      string
		length    byte
		strLength int
	}{
		{"Zero value", 0, 7},
		{"6", 6, 6},
		{"16", 16, 16},
	}

	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStr := GenPasswordAsBytes(tt.length)
			t.Log(string(gotStr))
			if len(gotStr) != tt.strLength {
				t.Errorf("GenPasswordAsBytes() got length = %v, want length %v", len(gotStr), tt.strLength)
			}
		})
	}
}
