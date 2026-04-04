/*

Name: Max/Min Value (Generics)
Level: Easy
TC: O(n)
Desc: Find the minimum and maximum value using generics
Note: cmp.Ordered allows to use <, >, <=, >= operators on the type K. It includes all built-in ordered types (integers, floats, strings) and any user-defined types that implement the necessary comparison operators.
*/

package main

import (
	"fmt"

	"cmp"
)

func minmax[K cmp.Ordered](a []K) (K, K) {
	var min, max K
	if len(a) > 0 {
		min = a[0]
		max = a[0]
	}

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
	fmt.Println(minmax([]string{"A", "B", "C"}))
	fmt.Println(minmax([]string{"a", "b", "c"}))
	fmt.Println(minmax([]float64{1.1416, 2.718, 3.1416}))
}
