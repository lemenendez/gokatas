//go:build ignore
// +build ignore

/*
Name: Linear Is Sorted
Level: Easy
TC: O(n)
SC: O(1)
Desc: Check whether a slice is already sorted by scanning adjacent elements.

By convention, a plain "isSorted" check usually means ascending order.
This kata also includes a descending version for learning and comparison.
*/

package main

import "fmt"

// isSortedAsc reports whether the slice is sorted in ascending order.
//
// This is the conventional introductory version of the algorithm: compare each
// element with the next one and stop as soon as the order is violated.
func isSortedAsc(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}

// isSortedDesc reports whether the slice is sorted in descending order.
//
// This is a useful variation of the same linear-scan idea, but in practice the
// default meaning of "sorted" usually refers to ascending order unless stated otherwise.
func isSortedDesc(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] < a[i+1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("ascending:")
	fmt.Println(isSortedAsc([]int{}))             // true: empty slice
	fmt.Println(isSortedAsc([]int{5}))            // true: one element
	fmt.Println(isSortedAsc([]int{1, 2, 3}))      // true
	fmt.Println(isSortedAsc([]int{1, 1, 1}))      // true: duplicates allowed
	fmt.Println(isSortedAsc([]int{-5, -2, 0, 7})) // true: includes negatives
	fmt.Println(isSortedAsc([]int{3, 2, 1}))      // false
	fmt.Println(isSortedAsc([]int{-1, -3, 2}))    // false

	fmt.Println("descending:")
	fmt.Println(isSortedDesc([]int{}))             // true: empty slice
	fmt.Println(isSortedDesc([]int{5}))            // true: one element
	fmt.Println(isSortedDesc([]int{3, 2, 1}))      // true
	fmt.Println(isSortedDesc([]int{4, 4, 2}))      // true: duplicates allowed
	fmt.Println(isSortedDesc([]int{7, 0, -2, -5})) // true: includes negatives
	fmt.Println(isSortedDesc([]int{1, 2, 3}))      // false
	fmt.Println(isSortedDesc([]int{-3, -1, -2}))   // false
}
