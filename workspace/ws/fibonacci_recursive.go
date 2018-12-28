package ws

func FibonacciRecursive(n int) (fib []int) {
	for i := 1; i <= n; i++ {
		fib = append(fib, rfibonacci(i))
	}
	return
}

func rfibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return rfibonacci(n-1) + rfibonacci(n-2)
	}
}
