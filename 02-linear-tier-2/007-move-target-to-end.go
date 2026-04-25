//go:build ignore
// +build ignore

package main

import "fmt"

/*
Name: Linear Move Any Target To End (Placeholder)
Level: Easy
TC: O(n)
SC: O(1) extra
Desc: Generalize move-zeros by moving any target value k to the end.
*/

// TODO:
// 1) Reuse the same compaction pattern
// 2) Keep values where a[i] != target
// 3) Fill tail with target value
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
}
