//go:build ignore
// +build ignore

/*
Name: Memoization Fibonacci (dynamic programming)
Level: Medium
TC: O(n)
SP: O(n)
*/

package main

import (
	"fmt"
)

func fib(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}
	if val, ok := memo[n]; ok {
		return val
	}
	memo[n] = fib(n-1, memo) + fib(n-2, memo)
	return memo[n]
}

func main() {
	memo := make(map[int]int, 2)
	memo[0] = 0
	memo[1] = 0
	fmt.Println(fib(15, memo))

}
