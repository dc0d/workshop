package primefactors

func Generate(n int) (result []int) {
	for ; n > 2 && n%2 == 0; n /= 2 {
		result = append(result, 2)
	}
	if n > 1 {
		result = append(result, n)
	}
	return
}
