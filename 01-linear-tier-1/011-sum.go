//go:build ignore
// +build ignore

/*
Name: Linear Sum
Level: Easy
TC: O(n)
SC: O(1)
Desc: Compute the sum of all values in a slice using a single linear scan.

This is a classic introductory algorithm because it demonstrates accumulation
with a running total and handles empty input naturally.
*/

package main

import "fmt"

// sum returns the total of all integers in the slice.
//
// It starts from 0 and adds each element once, making it a simple O(n)
// accumulation algorithm.
func sum(a []int) int {
	res := 0
	for i := range a {
		res += a[i]
	}
	return res
}

func main() {
	fmt.Println(sum([]int{}))              // 0: empty slice
	fmt.Println(sum([]int{5}))             // 5: one element
	fmt.Println(sum([]int{1, 2, 3}))       // 6
	fmt.Println(sum([]int{0, 0, 0}))       // 0: all zeros
	fmt.Println(sum([]int{-1, -2, -3}))    // -6: negatives only
	fmt.Println(sum([]int{-5, 10, -2}))    // 3: mixed positive and negative
	fmt.Println(sum([]int{100, 200, 300})) // 600
}
