package main

import "github.com/01-edu/z01"

func setPoint(x, y *string) {
	*x = "42"
	*y = "21"
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func main() {
	var firstring string
	var secondstring string

	setPoint(&firstring, &secondstring)

	printStr("x = " + firstring + ", y = " + secondstring)
}
