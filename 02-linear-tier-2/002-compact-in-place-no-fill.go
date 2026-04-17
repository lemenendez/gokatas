//go:build ignore
// +build ignore

package main

import "fmt"

/*
Name: Linear Compact In Place Without Fill
Level: Easy
TC: O(n)
SC: O(1) extra
Desc: Compact non-zeros at the front and return the logical length.
*/

func compactInPlaceNoFill(a []int) int {
	write := 0

	for i := range a {
		if a[i] != 0 {
			a[write] = a[i]
			write++
		}
	}
	return write
}

func main() {
	a := []int{1, 0, 3, 0, 4}
	n := compactInPlaceNoFill(a)
	fmt.Println(a[:n]) // [1 3 4]

	b := []int{0, 0, 0}
	n = compactInPlaceNoFill(b)
	fmt.Println(b[:n]) // []

	c := []int{1, 2, 3}
	n = compactInPlaceNoFill(c)
	fmt.Println(c[:n]) // [1 2 3]

	d := []int{}
	n = compactInPlaceNoFill(d)
	fmt.Println(d[:n]) // []
}
