//go:build ignore
// +build ignore

/*

Name: Maximum Value
Source: Graphic Go Algorithms Book
Credits: Yan Hu
Level: Easy
TC: O(n)
Desc: Compare and swap elements in the array until the last number in the array is the maximum

*/

package main

import "fmt"

func max(a []int, length int) int {
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
	var length = len(values)
	fmt.Println("Before Swap:", values)
	var maxVal = max(values, length)
	fmt.Printf("Max value = %d\n", maxVal)
	fmt.Println("After Swap:", values)
}
