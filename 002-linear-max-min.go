/*

Name: Max/Min Value
Level: Easy
TC: O(n)
Desc: Find the minium and max avalue

*/

package main

import (
	"fmt"
	"math"
)

func minmax(a []int) (int, int) {
	min := math.MaxInt
	max := math.MinInt

	for i := 0; i < len(a); i++ {
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

}
