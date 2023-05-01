package piscine

func Swap(a *int, b *int) {
	var c int
	c = *b
	*b = *a
	*a = c
}
