package piscine

func checkup(a rune) bool {
	if a >= 'A' && a <= 'Z' || a >= 'a' && a <= 'z' || a >= '0' && a <= '9' {
		return true
	}
	return false
}

func IsAlpha(s string) bool {
	runes := []rune(s)

	for i := range runes {
		if checkup(runes[i]) == false {
			return false
		}
	}
	return true
}
