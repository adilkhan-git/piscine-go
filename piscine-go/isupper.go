package piscine

func check(a rune) bool {
	if a >= 'A' && a <= 'Z' {
		return true
	}
	return false
}

func IsUpper(s string) bool {
	runes := []rune(s)

	for i := range runes {
		if check(runes[i]) == false {
			return false
		}
	}
	return true
}
