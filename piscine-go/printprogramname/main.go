package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	pn := arguments[0]
	for b, a := range pn {
		if b >= 2 {
			z01.PrintRune(a)
		}
	}
	z01.PrintRune('\n')
}
