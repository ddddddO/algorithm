package ws

// 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, ・・・
// targetは0始まり
func Fibonacci(target int) (fib []int) {
	fib = []int{1, 1}

	switch target {
	case 0:
		fib = fib[:1]
		return
	case 1:
		return
	}

	sum := 0
	for i := 1; i < target; i++ {
		sum = fib[i-1] + fib[i]
		fib = append(fib, sum)
	}
	return
}
