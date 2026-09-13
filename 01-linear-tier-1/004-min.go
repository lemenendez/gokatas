//go:build ignore
// +build ignore

/*

Name: Minimum Value
TC: O(n)
Desc: Finds the minimum value in an array of integers.
returns the minimum value in the array.


*/

package main

import "fmt"

func min(arr []int, length int) int {
	minIdx := 0
	for i := 0; i < length; i++ {
		if arr[minIdx] > arr[i] {
			minIdx = i
		}
	}
	return arr[minIdx]
}

func main() {
	scores := []int{60, 50, 95, 80, 70}
	length := len(scores)
	min := min(scores, length)
	fmt.Printf("Min Value=%d\n", min)
}
