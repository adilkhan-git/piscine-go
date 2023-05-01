package piscine

func StrRev(s string) string {
	Astr := []rune(s)
	count := 0
	for i := range Astr {
		i++
		count++
	}

	var reverse string
	for i := count - 1; i >= 0; i-- {
		reverse += string(Astr[i])
	}
	return reverse
}
