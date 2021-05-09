package assets

// SetByte checks and sets input string to a given pointer to byte type.
func SetByte(inputByte *byte, inputString string) bool {

	var err error
	*inputByte, err = String2Byte(inputString)
	return err != nil
}

// SetUint32 checks and sets input string to a given pointer to uint16 type.
func SetUint32(inputUint32 *uint32, inputString string) bool {

	var err error
	*inputUint32, err = String2Uint32(inputString)
	return err != nil
}
