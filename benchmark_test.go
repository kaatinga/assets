package assets

import (
	"strconv"
	"testing"
)

func BenchmarkUint16String(b *testing.B) {

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Uint16String(199)
		Uint16String(19999)
		Uint16String(55)
	}
}

func BenchmarkStrvconvItoa(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		strconv.Itoa(199)
		strconv.Itoa(19999)
		strconv.Itoa(55)
	}
}
