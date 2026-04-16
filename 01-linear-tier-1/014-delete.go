//go:build ignore
// +build ignore

/*
Name: Linear Delete
Level: Easy
TC: O(n)
SC: O(n) for the manual version, O(1) extra for the in-place version
Desc: Delete a value from a slice at a given index.

This kata shows the same idea implemented in two styles:
1. a manual copy version,
2. an in-place version.
*/

package main

import "fmt"

// deleteManual removes the element at index by creating a new slice.
//
// This version is easy to understand for practice because it copies the left
// part first and then copies the right part one position to the left.
func deleteManual(a []int, index int) []int {
	if index < 0 || index >= len(a) {
		panic("out of range")
	}
	l := len(a)
	temp := make([]int, l-1)
	for i := 0; i < index; i++ {
		temp[i] = a[i]
	}
	for i := index + 1; i < l; i++ {
		temp[i-1] = a[i]
	}
	return temp
}

// deleteInplace removes the element at index by shifting values left in the
// same underlying slice.
//
// This version uses less extra memory because it does not allocate a second
// slice; it overwrites the removed position and returns a shorter view.
func deleteInplace(a []int, index int) []int {
	if index < 0 || index >= len(a) {
		panic("out of range")
	}
	lim := len(a) - 1
	// left shift all
	for i := index; i < lim; i++ {
		a[i] = a[i+1]
	}
	return a[:lim]
}

func main() {
	a := []int{1, 2, 3}
	a = deleteManual(a, 1)
	fmt.Println(a)

	b := []int{2, 4, 6, 8}
	b = deleteManual(b, 2)
	fmt.Println(b)

	c := []int{10, 20, 30, 40}
	c = deleteManual(c, 1)
	fmt.Println(c)

	a = []int{1, 2, 3}
	a = deleteInplace(a, 1)
	fmt.Println(a)

	b = []int{2, 4, 6, 8}
	b = deleteInplace(b, 2)
	fmt.Println(b)

	c = []int{10, 20, 30, 40}
	c = deleteInplace(c, 1)
	fmt.Println(c)

}
