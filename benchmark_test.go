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

// nolint
func BenchmarkString2Byte(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		String2Byte("0")
		String2Byte("255")
		String2Byte("25549672951")
	}
}

// nolint
func BenchmarkString2Uint32(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		String2Uint32("0")
		String2Uint32("255")
		String2Uint32("25549672951")
	}
}

// nolint
func BenchmarkString2Uint16(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		String2Uint16("0")
		String2Uint16("255")
		String2Uint16("25549672951")
	}
}

// nolint
func BenchmarkStrvconv_Atoi(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		strconv.Atoi("0")
		strconv.Atoi("255")
		strconv.Atoi("25549672951")
	}
}

//func BenchmarkBytes2Uint32(b *testing.B) {
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		Bytes2Uint32([]byte{0})
//		Bytes2Uint32([]byte{0x32, 0x35, 0x35})
//		Bytes2Uint32([]byte{0x32, 0x35, 0x35, 0x34, 0x39, 0x36, 0x37, 0x32, 0x39, 0x35, 0x31})
//	}
//}

//func BenchmarkStrvconv_AtoiPlusBytesConversion(b *testing.B) {
//	b.ReportAllocs()
//	for i := 0; i < b.N; i++ {
//		strconv.Atoi(string([]byte{0}))
//		strconv.Atoi(string([]byte{0x32, 0x35, 0x35})
//		//strconv.Atoi(string([]byte{0x36, 0x36, 0x35, 0x35, 0x35}))
//	}
//}
