package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printstr(s string) {
	srune := []rune(s)
	for i := 0; i < len(srune); i++ {
		z01.PrintRune(srune[i])
	}
	z01.PrintRune('\n')
}

func main() {
	for i := len(os.Args) - 1; i > 0; i-- {
		printstr(os.Args[i])
	}
}
