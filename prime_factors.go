package primefactors

func Generate(n int) (result []int) {
	if n > 1 {
		result = append(result, n)
	}
	return
}
