package piscine

func Swap(a *int, b *int) {
	*a, *b, *a = *b, *a, *b
}
