package controllers

func fibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	} else {
		return fibonacci(n-2) + fibonacci(n-1)
	}
}

// GetFibSequence returns the sequence of values
func GetFibSequence(n int) (values []int) {
	done := make(chan interface{})
	defer close(done)
	fn := fibonacci
	for num := range take(done, repeatFunc(done, n, fn), n) {
		values = append(values, num.(int))
	}
	return values
}
