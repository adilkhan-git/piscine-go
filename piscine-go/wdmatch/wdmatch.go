package main

import (
	"os"

	"github.com/01-edu/z01"
)

func WdMatch(s1 string, s2 string) bool {
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
		if WdMatch(os.Args[1], os.Args[2]) {
			for _, j := range os.Args[1] {
				z01.PrintRune(j)
			}
			z01.PrintRune('\n')
		}
	}
}
