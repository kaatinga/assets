package assets

// Deprecated: SuperBytesToUint32 checks and converts input string to uint32 type
func SuperBytesToUint32(input string) (output uint32, ok bool) {

	var output64 uint64

	if input[0] < 48 || input[0] > 57 {
		return
	}

	var (
		value rune
		key   int
	)

MainLoop:
	for key, value = range input {

		switch key {
		case 0:
			// processing the first byte
			output64 = uint64(value) - 48
			continue MainLoop
		default:
			if value < 48 || value > 57 {
				ok = true
				break MainLoop
			}
		}

		// multiply to 10 and sum
		output64 = output64*10 + uint64(value) - 48

		if output64 >= maxUint32 {
			// the number exceeds the uint32 limit!
			return 0, false
		}
	}

	output = uint32(output64)
	ok = true
	return
}
