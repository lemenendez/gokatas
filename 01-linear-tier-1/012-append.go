//go:build ignore
// +build ignore

/*
Name: Linear Append
Level: Easy
TC: O(n)
SC: O(n)
Desc: Appends a new value at the end of the array. https://go.dev/blog/slices
Allocates space for the current array plus new value, place new value at the end
Do not use built-in Append function, we can use copy however to copy elements from the old array to the new one.

*/

package main

import "fmt"

func appendLinear(a []int, b int) []int {
	l := len(a)
	temp := make([]int, l+1)
	/*
		another way to copy elements from the old array to the new one is using a for loop:
		for i := range a {
			temp[i] = a[i]
		}
	*/
	// go recomendation is to use the built-in copy function
	copy(temp, a)
	temp[l] = b
	return temp
}

func main() {
	a := []int{1, 2}
	a = appendLinear(a, 3)
	fmt.Println(a)
}
