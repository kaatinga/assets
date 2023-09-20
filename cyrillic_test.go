package assets

import "testing"

func TestIsCyrillicString(t *testing.T) {
	tests := []struct {
		name    string
		company string
		wantOk  bool
	}{
		{`string1`, "ООО «аб_в»", false},
		{`string+digits`, "ООО «1а2б_3в»", false},
		{`string2`, "ООО «Про&+ба»", true},
		{`string+digits2`, "ООО «а-б1-в»", true},
		{`string3`, `ООО "а&бв"`, true},
		{`english string`, "ООО «Company»", false},
		{`english string`, "ООО «Ромашка №1»", true},
		{`some chars`, `&"+-»«.,/№ текст`, true},
	}

	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOk := IsCyrillicString(tt.company); gotOk != tt.wantOk {
				t.Errorf("CheckCompanyName() = %v, want %v", gotOk, tt.wantOk)
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
		{"ok", args{"Кирилица текст"}, true},
		{"english", args{"Cyrillic String"}, false},
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
