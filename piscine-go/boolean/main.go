package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	arrayStr := []rune(s)
	for _, a := range arrayStr {
		z01.PrintRune(a)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 1 {
		return false
	} else {
		return true
	}
}

func main() {
	arguments := os.Args
	lengthOfArg := 0
	for range arguments {
		lengthOfArg++
	}
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	if isEven(lengthOfArg + 1) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
