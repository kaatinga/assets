package assets

var (
	// Deprecated: SuperBytesToUint32 is an alias for Bytes2Uint32.
	SuperBytesToUint32 = Bytes2Uint32
)

// Deprecated: StByte checks and converts input string to byte type. You had better use String2Byte().
func StByte(inputString string) (byte, bool) {
	var (
		tmpUint64 uint64 // a temporary int value
		ok        bool
	)

	tmpUint64, ok = StUint64(inputString)
	if ok && tmpUint64 <= maxUint8 {
		return byte(tmpUint64), true
	}

	return 0, false
}

// Deprecated: StUint32 checks and converts input string to uint32 type. You had better use String2Uint32().
func StUint32(inputString string) (uint32, bool) {
	var (
		tmpUint64 uint64 // a temporary int value
		ok        bool
	)

	tmpUint64, ok = StUint64(inputString)
	if ok && tmpUint64 <= maxUint32 {
		return uint32(tmpUint64), true
	}

	return 0, false
}