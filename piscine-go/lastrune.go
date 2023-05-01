package piscine

func LastRune(s string) rune {
	sentence := []rune(s)

	length := 0

	for i := range sentence {
		length = i + 1
	}
	return rune(sentence[length-1])
}
