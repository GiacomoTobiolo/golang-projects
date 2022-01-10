package piscine

func Any(f func(string) bool, a []string) bool {
	numeric := false
	for i := 0; i < len(a); i++ {
		if f(a[i]) {
			return true
		} else if !f(a[i]) && !numeric {
			numeric = false
		}
	}
	return numeric
}
