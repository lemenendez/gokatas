package main

import "fmt"

func removeDupsInvariant(a []int) []int {
	writer := 1
	for i := 1; i < len(a); i++ {
		if a[i] != a[writer-1] {
			a[writer], a[i] = a[i], a[writer]
			writer++
		}
	}
	return a[0:writer]
}

func main() {
	fmt.Println(removeDupsInvariant([]int{1, 1}))
	fmt.Println(removeDupsInvariant([]int{1, 2, 2}))
	fmt.Println(removeDupsInvariant([]int{1, 1, 2, 2}))
	fmt.Println(removeDupsInvariant([]int{1, 1, 1}))
	fmt.Println(removeDupsInvariant([]int{1, 2, 3, 4, 5, 5, 6, 6}))
	fmt.Println(removeDupsInvariant([]int{5, 5, 6, 6}))
}
