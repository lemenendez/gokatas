//go:build ignore
// +build ignore

package main

import "fmt"

/*
Name: Linear Move Any Target To End
Level: Easy
TC: O(n)
SC: O(1) extra
Desc: Generalize move-zeros by moving any target value k to the end.
Use compaction pattern with swap + write pointer.
*/

// moveTargetToEnd moves every target value to the end in-place.
// It preserves the relative order of non-target values.
// TC: O(n)
// SC: O(1) extra
func moveTargetToEnd(a []int, target int) []int {
	writer := 0
	for i := 0; i < len(a); i++ {
		if a[i] != target {
			a[writer], a[i] = a[i], a[writer]
			writer++
		}
	}
	return a
}

func main() {
	fmt.Println(moveTargetToEnd([]int{1, 0, 2, 0, 3}, 0))
	fmt.Println(moveTargetToEnd([]int{0, 0, 1, 2, 3}, 0))
	fmt.Println(moveTargetToEnd([]int{1, 2, 3}, 0))
	fmt.Println(moveTargetToEnd([]int{7, 7, 7}, 7))
	fmt.Println(moveTargetToEnd([]int{}, 0))
	fmt.Println(moveTargetToEnd([]int{2, 1, 2, 3, 2, 4}, 2))
}
