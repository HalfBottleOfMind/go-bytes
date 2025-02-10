package gobytes

func BoolsToBytes(input []bool) []byte {
	l := len(input) / 8
	if len(input)%8 > 0 {
		l += 1
	}

	result := make([]byte, l)
	for i := range l {
		start := i * 8
		eightBits := make([]bool, 8)
		for j := range 8 {
			b := false
			if len(input) > start+j {
				b = input[start+j]
			}
			eightBits[j] = b
		}
		result[i] = BoolsToByte([8]bool(eightBits))
	}

	return result
}

func BoolsToByte(input [8]bool) (out byte) {
	if len(input) != 8 {
		panic("input length should be exactly 8")
	}

	for i := range 8 {
		if input[i] {
			out += 1 << (7 - i)
		}
	}

	return
}
