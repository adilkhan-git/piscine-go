package main

import (
	"os"

	"github.com/01-edu/z01"
)

func HiddenP(s1 string, s2 string) bool {
	runes1 := []rune(s1)
	runes2 := []rune(s2)
	var rest string
	count := 0
	for i := 0; i < len(runes1); i++ {
		for j := count; j < len(runes2); j++ {
			if runes1[i] == runes2[j] {
				rest += string(runes1[i])
				break
			}
			count++
		}
	}
	return s1 == rest
}

func main() {
	if len(os.Args) == 3 {
		if HiddenP(os.Args[1], os.Args[2]) {
			z01.PrintRune('1')
		} else if !HiddenP(os.Args[1], os.Args[2]) {
			z01.PrintRune('0')
		}
	}
	z01.PrintRune('\n')
}
