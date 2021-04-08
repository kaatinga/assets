package assets

import (
	"strconv"
	"testing"
)

//func BenchmarkUint162Bytes(b *testing.B) {
//
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		Uint162Bytes(199)
//		Uint162Bytes(0)
//		Uint162Bytes(55)
//	}
//}
//
//func BenchmarkByte2Bytes(b *testing.B) {
//
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		Byte2Bytes(199)
//		Byte2Bytes(0)
//		Byte2Bytes(55)
//	}
//}

//func BenchmarkByte2String(b *testing.B) {
//
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		Byte2String(199)
//		Byte2String(0)
//		Byte2String(55)
//	}
//}

func BenchmarkUint162String(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Uint162String(19999)
		Uint162String(0)
		Uint162String(55)
	}
}

func BenchmarkStrvconvItoa(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		strconv.Itoa(19999)
		strconv.Itoa(0)
		strconv.Itoa(55)
	}
}
