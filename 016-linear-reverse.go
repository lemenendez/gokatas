//go:build ignore
// +build ignore

/*
Name: Linear Reverse
Level: Easy
TC: O(n)
SC: O(n) for reverse, O(1) extra for reversev2 (in-place)
Desc: Reverse the order of a slice using linear-time strategies.

This kata shows two common versions:
1. out-of-place reverse that creates a new slice,
2. in-place reverse using two pointers and swaps up to the middle.
*/

package main

import "fmt"

// reverse returns a new slice with elements in reverse order.
//
// This version keeps the input unchanged and writes into a new output slice.
func reverse(a []int) []int {
	l := len(a)
	b := make([]int, l)
	for i, j := 0, l-1; i < l; i, j = i+1, j-1 {
		b[i] = a[j]
	}
	return b
}

// reversev2 reverses the slice in place using two pointers.
//
// Only half of the positions need swaps: after reaching the middle, every
// remaining pair has already been swapped from the opposite side.
func reversev2(a []int) []int {
	// reverse inplace
	l := len(a)
	mid := len(a) / 2
	for i, j := 0, l-1; i < mid; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}

func main() {
	fmt.Println("reverse (new slice):")
	fmt.Println(reverse([]int{}))             // []
	fmt.Println(reverse([]int{5}))            // [5]
	fmt.Println(reverse([]int{1, 2, 3}))      // [3 2 1]
	fmt.Println(reverse([]int{2, 2, 2}))      // [2 2 2]
	fmt.Println(reverse([]int{-5, 0, 7, -1})) // [-1 7 0 -5]

	fmt.Println("reversev2 (in-place):")
	b := []int{100, 1200, 300}
	fmt.Println(reversev2(b)) // [300 1200 100]
	fmt.Println(b)            // same underlying slice changed
	fmt.Println(reversev2([]int{}))
	fmt.Println(reversev2([]int{9}))
	fmt.Println(reversev2([]int{1, 2, 3, 4}))
	fmt.Println(reversev2([]int{-3, -2, -1, 0, 1}))
}
