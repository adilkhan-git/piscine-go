package piscine

func NRune(s string, n int) rune {
	if n <= 0 {
		return 0
	}

	sentence := []rune(s)

	length := 0

	for i := range sentence {
		length = i + 1
	}
	if n <= length {
		return rune(sentence[n-1])
	} else {
		return 0
	}
}
