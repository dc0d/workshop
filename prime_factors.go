package primefactors

func Generate(n int) (result []int) {
	for p := 2; n > p; p++ {
		for ; n%p == 0; n /= p {
			result = append(result, p)
		}
	}
	if n > 1 {
		result = append(result, n)
	}
	return
}
