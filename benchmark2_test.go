package assets

import "testing"

func BenchmarkGenPassword(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		GenPassword(40)
	}
}

func BenchmarkGenPasswordAsBytes(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		GenPasswordAsBytes(40)
	}
}
