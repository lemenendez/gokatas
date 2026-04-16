//go:build ignore
// +build ignore

/*
Name: Linear Count Evens and Odds
Level: Easy
TC: O(n)
SC: O(1)
Desc: Count how many even and odd values exist in a slice using a single pass.

This kata is a good introductory example of classification while iterating,
where each element is checked once and added to one of two counters.
*/

package main

import "fmt"

// countEvensOdds returns two counters: the number of even values and the
// number of odd values in the input slice.
//
// It scans the slice once and classifies each element using the remainder
// operator.
func countEvensOdds(a []int) (int, int) {
	evens := 0
	odds := 0
	for i := range a {
		if a[i]%2 == 0 {
			evens += 1
		} else {
			odds += 1
		}
	}
	return evens, odds
}

func main() {
	fmt.Println(countEvensOdds([]int{}))              // 0 evens, 0 odds
	fmt.Println(countEvensOdds([]int{1, 2, 3}))       // 1 even, 2 odds
	fmt.Println(countEvensOdds([]int{1, 1, 1}))       // 0 evens, 3 odds
	fmt.Println(countEvensOdds([]int{2, 2, 2}))       // 3 evens, 0 odds
	fmt.Println(countEvensOdds([]int{0, 4, 7, 9}))    // 2 evens, 2 odds
	fmt.Println(countEvensOdds([]int{-2, -1, -4, 5})) // 2 evens, 2 odds
	fmt.Println(countEvensOdds([]int{10}))            // 1 even, 0 odds
	fmt.Println(countEvensOdds([]int{11}))            // 0 evens, 1 odd
}
