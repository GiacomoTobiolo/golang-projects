package piscine

func Map(f func(int) bool, a []int) []bool {
	boolarr := make([]bool, len(a))
	for i := 0; i < len(a); i++ {
		boolarr[i] = f(a[i])
	}

	return boolarr
}
