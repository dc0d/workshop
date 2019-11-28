package workshop

func Fib(n int) (fib int) {
	for a, b := 1, 1; n > 0; a, b, n = b, a+b, n-1 {
		fib = a
	}

	return
}
