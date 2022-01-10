package piscine

func Index(s string, toFind string) int {
	aSlice := []rune(s)
	finder := []rune(toFind)
	if s == "" || toFind == "" {
		return 0
	}
	if (len(s) == len(toFind)) && (s == toFind) {
		return 0
	}
	j := 0
	for i := 0; i <= (len(s) - 1); i++ {
		if aSlice[i] == finder[j] {
			j = j + 1
			if j == len(toFind) {
				return (i - len(toFind) + 1)
			}
		} else {
			j = 0
		}
	}
	return -1
}
