//go:build ignore
// +build ignore

/*
Name: Linear Count (Generics)
Level: Easy
TC: O(n)
Desc: Count the occurrences of a target element in an array using linear search
*/

package main

import (
	"fmt"
)

func count[K comparable](arr []K, target K) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println("count:", count([]int{1, 2, 3}, (3)))
	fmt.Println("count:", count([]int{1, 2, 3, 4, 5, 3}, (3)))
	fmt.Println("count:", count([]int{1, 2, 5}, (3)))

	fmt.Println("count:", count([]string{"A", "B", "C"}, ("C")))
}
