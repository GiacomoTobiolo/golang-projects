package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	nameaprogramme := os.Args[0]
	nameaprogrammerune := []rune(nameaprogramme)

	for index, s := range nameaprogrammerune {
		if index > 1 {
			z01.PrintRune(s)
		}
	}
	z01.PrintRune('\n')
}
