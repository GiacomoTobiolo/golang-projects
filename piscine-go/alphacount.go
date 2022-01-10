package piscine

func AlphaCount(s string) int {
	// aRune := []rune(s)
	counter := 0

	for _, char := range s {
		if (char >= 65 && char <= 90) || (char >= 97 && char <= 122) {
			counter++
		}
	}
	return counter
}
