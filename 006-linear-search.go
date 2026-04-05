/*

Name: Linear Search (Generics)
Level: Easy
TC: O(n)
Desc: Find the index of a target element in an array using linear search
*/

package main

import "fmt"

func linearSearch(arr []int, target int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println("index:", linearSearch([]int{1, 2, 3}, (3)))
	fmt.Println("index:", linearSearch([]int{1, 2, 3}, (100)))
}
