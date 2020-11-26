package Fibonacci

func Fibonacci(limit int) []int {
	fibonacciList := []int{0, 1}
	for i := 1; i < limit-1; i++ {
		n := fibonacciList[i] + fibonacciList[i-1]
		fibonacciList = append(fibonacciList, n)
	}

	return fibonacciList
}
