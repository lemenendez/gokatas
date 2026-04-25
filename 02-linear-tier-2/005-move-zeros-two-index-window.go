//go:build ignore
// +build ignore

package main

import "fmt"

/*
Name: Linear Move Zeros Right (Two-Index Window)
Level: Easy
TC: O(n)
SC: O(1)
Desc: Move zeros right in-place using explicit left/right indexes.
*/

func moveZerosRightTwoIndexWindow(a []int) []int {
	left := 0
	for right := 0; right < len(a); right++ {
		if a[right] != 0 {
			a[left], a[right] = a[right], a[left]
			left++
		}
	}
	return a
}

func main() {
	fmt.Println(moveZerosRightTwoIndexWindow([]int{1, 0, 2, 0, 3}))
}
