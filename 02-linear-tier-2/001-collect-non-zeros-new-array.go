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

func collectNonZeros(a []int) []int {
	result := make([]int, 0, len(a))
	for i := range a {
		if a[i] != 0 {
			result = append(result, a[i])
		}
	}
	return result
}

func main() {

	fmt.Println(collectNonZeros([]int{1, 3, 0, 4}))

}
