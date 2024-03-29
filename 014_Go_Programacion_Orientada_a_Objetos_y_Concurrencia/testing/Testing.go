package testing

import "time"

func Sum(x, y int) int {
	return x + y
}

func Fibonacci(n int) int {
	fib := make([]int, n+1)
	fib[0] = 0
	fib[1] = 1
	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
		time.Sleep(time.Duration(50) * time.Microsecond)
	}
	return fib[n]
}
