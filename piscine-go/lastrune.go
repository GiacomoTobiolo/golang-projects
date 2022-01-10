package piscine

func LastRune(s string) rune {
	sRune := []rune(s)
	y := len(s)
	return sRune[y-1]
}
