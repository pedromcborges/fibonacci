package main

import (
	"fmt"
)

func main() {
	var n int
	_, _ = fmt.Scanf("%d", &n)

	memo := make([]uint, n+1)

	result := fibonacci(n, memo)
	fmt.Println(result)
}

func fibonacci(n int, memo []uint) uint {
	if memo[n] != 0 {
		return memo[n]
	}

	if n == 1 || n == 2 {
		memo[n] = 1
	} else {
		memo[n] = fibonacci(n-1, memo) + fibonacci(n-2, memo)
	}
	fmt.Println(memo)
	return memo[n]
}
