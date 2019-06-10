package modify

func Swap(a, b int) (int, int) {
	return b, a
}

func Swap2(a, b *int) {
	*a, *b = *b, *a
}
