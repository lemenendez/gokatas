//go:build ignore
// +build ignore

package main

import "fmt"

/*
Name: Linear Collect Non-Zeros Into New Array
Level: Easy
TC: O(n)
SC: O(n)
Desc: Build a new array with non-zero values in the same order.
*/

// SC: O(n) — allocates a second array to hold the non-zero values.
// This is the simplest approach but NOT space-optimal: in the worst case
// (no zeros) result is the same size as the input.
// The in-place variants (002, 003, ...) solve this with SC: O(1).
func collectNonZeros(a []int) []int {
	// Pre-allocate with cap=len(a) to avoid reallocations, but length starts at 0.
	result := make([]int, 0, len(a))
	for i := range a {
		if a[i] != 0 {
			result = append(result, a[i]) // copy non-zero into the new array
		}
		// zeros are simply skipped — they never enter result
	}
	return result
}

func main() {

	fmt.Println(collectNonZeros([]int{1, 3, 0, 4}))

}
