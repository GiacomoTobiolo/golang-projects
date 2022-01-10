package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	var mySlice []int
	var digit int = 0
	for n != 0 {
		digit = n % 10
		mySlice = append(mySlice, digit)
		n /= 10
	}

	for i := '0'; i <= '9'; i++ {
		for _, character := range mySlice {
			if i == rune(character+48) {
				z01.PrintRune(i)
			}
		}
	}
}
