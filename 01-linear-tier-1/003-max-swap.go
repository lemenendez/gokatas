//go:build ignore
// +build ignore

/*

Name: Find Max
Level: Easy
TC: O(n)
Desc: Find max value by swapping, max value is at the end of the array after

*/

package main

import "fmt"

func max(a []int) int {
	length := len(a)
	if length == 0 {
		return -1
	}
	for i := 0; i < length-1; i++ {
		if a[i] > a[i+1] {
			// Usual swap mechanism
			// 	var temp = a[i]
			// 	a[i] = a[i+1]
			// 	a[i+1] = temp
			// Go multi assigment mechanism
			a[i], a[i+1] = a[i+1], a[i]
		}
	}
	max := a[length-1]
	return max
}

func main() {
	var values = []int{60, 50, 95, 80, 70}
	fmt.Println("Before:", values)
	var maxVal = max(values)
	fmt.Printf("Max value = %d\n", maxVal)
	fmt.Println("After:", values)
}
