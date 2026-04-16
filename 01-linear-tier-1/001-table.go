//go:build ignore
// +build ignore

/*
Name: Linear Table Definition
Source: Graphic Go Algorithms Book
Credits: Yan Hu
Level: Easy
TC: O(n)
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

func main() {
	var scores = []int{90, 70, 50, 80, 60, 85}

	var length = len(scores)

	print(scores, length)
}
