//go:build ignore
// +build ignore

/*
Name: Linear Append
Level: Easy
TC: O(n)
SC: O(n)
Desc: Appends a new value at the end of the array. https://go.dev/blog/slices
Allocates space for the current array plus new value, place new value at the end


*/

package main

import "fmt"

func appendLinear(a []int, b int) []int {
	l := len(a)
	temp := make([]int, l+1)
	for i := range a {
		temp[i] = a[i]
	}
	temp[l] = b
	return temp
}

func main() {
	a := []int{1, 2}
	a = appendLinear(a, 3)
	fmt.Println(a)
}
