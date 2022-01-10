package piscine

func IsAlpha(s string) bool {
	if s == "" {
		return true
	}
	for _, i := range s {
		if i < 48 {
			return false
		} else if i > 57 && i < 65 {
			return false
		} else if i > 90 && i < 97 {
			return false
		} else if i > 122 {
			return false
		}
	}
	return true
}
