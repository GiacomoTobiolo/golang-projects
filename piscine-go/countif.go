package piscine

func CountIf(f func(string) bool, tab []string) int {
	numeric := 0
	istrue := false
	for i := 0; i < len(tab); i++ {
		if f(tab[i]) {
			numeric++
			istrue = true
		} else if !f(tab[i]) && !istrue {
			istrue = false
		}
	}
	return numeric
}
