//go:build ignore
// +build ignore

package main

import "fmt"

/*
Name: Linear Move Zeros Left (Mirror Version)
Level: Easy
TC: O(n)
SC: O(1) extra
Desc: Mirror of move-zeros-right by reversing scan/write direction.
*/

func moveZerosLeftMirror(a []int) []int {
	writer := len(a) - 1
	for i := writer; i >= 0; i-- {
		if a[i] != 0 {
			a[writer], a[i] = a[i], a[writer]
			writer--
		}
	}

	return a
}

func main() {
	fmt.Println(moveZerosLeftMirror([]int{1, 0, 2, 0, 3}))
}
