package piscine

func IsPrintable(s string) bool {
	for _, i := range s {
		if i >= 7 && i <= 13 {
			return false
		}
	}
	return true
}
