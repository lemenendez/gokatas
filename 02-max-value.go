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

func print(a []int, count int) {
	for i := 0; i < count; i++ {
		if i == count-1 {
			fmt.Printf("%d\n", a[i])
		} else {
			fmt.Printf("%d, ", a[i])
		}
	}
}

func max(a []int, length int) int {
	for i := 0; i < length-1; i++ {
		if a[i] > a[i+1] {
			// Usual swap mechanism
			// var temp = a[i]
			// a[i] = a[i+1]
			// a[i+1] = temp
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
	fmt.Print("Before Swap:")
	print(values, length)
	var maxVal = max(values, length)
	fmt.Printf("Max value = %d\n", maxVal)
	fmt.Print("After Swap:")
	print(values, length)
}
