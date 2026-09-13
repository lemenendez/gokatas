//go:build ignore
// +build ignore

/*

Name: Max/Min Value
Level: Easy
TC: O(n)
Desc: Find both the minimum and maximum value of an unsorted array

*/

package main

import (
	"fmt"
	"math"
)

func minmax(a []int) (int, int) {
	min := math.MaxInt
	max := math.MinInt

	if len(a) == 0 {
		return min, max
	}

	for i := range a {
		if a[i] < min {
			min = a[i]
		}
		if a[i] > max {
			max = a[i]
		}
	}
	return min, max
}

func main() {

	fmt.Println(minmax([]int{1, 2, 3}))
	fmt.Println(minmax([]int{100, 200, -99}))
	fmt.Println(minmax([]int{1, 1, 1}))
	fmt.Println(minmax([]int{0, 0, 0, 0, 0}))
	fmt.Println(minmax([]int{1}))
	fmt.Println(minmax([]int{}))
}
