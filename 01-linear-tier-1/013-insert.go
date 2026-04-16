//go:build ignore
// +build ignore

/*
Name: Linear Insert
Level: Easy
TC: O(n)
SC: O(n)
Desc: Insert a value in the middle of a slice.

This kata shows the same idea implemented in three styles:
1. a classic two-loop version,
2. a single-loop version,
3. an idiomatic Go slice version.
*/

package main

import "fmt"

// insertMiddleGo inserts b at index len(a)/2 using Go's slice helpers.
//
// This is the most idiomatic version in Go because it combines copy and
// append to keep the code short and expressive.
func insertMiddleGo(a []int, b int) []int {
	l := len(a)
	middle := l / 2
	result := make([]int, middle+1)
	copy(result, a[0:middle])
	result[middle] = b
	result = append(result, a[middle:l]...)
	return result
}

// insertMiddleOneLoop inserts b at index len(a)/2 in a single pass.
//
// It is useful as a kata version because it makes the shifting logic explicit
// while still solving the problem with just one loop.
func insertMiddleOneLoop(a []int, b int) []int {
	l := len(a)
	middle := l / 2
	temp := make([]int, l+1)
	for i := 0; i < l+1; i++ {
		if i < middle {
			temp[i] = a[i]
			continue
		}
		if i == middle {
			temp[i] = b
			continue
		}
		temp[i] = a[i-1]
	}
	return temp
}

// insertMiddle inserts b at index len(a)/2 using two loops.
//
// This version is the most verbose, but it is great for learning because the
// copy-before / insert / copy-after steps are easy to follow.
func insertMiddle(a []int, b int) []int {
	l := len(a)
	middle := l / 2
	temp := make([]int, l+1)
	// copy from the beginning up to, but not including, the middle
	for i := 0; i < middle; i++ {
		temp[i] = a[i]
	}
	// set new value
	temp[middle] = b
	// copy the rest
	for i := middle + 1; i < l+1; i++ {
		temp[i] = a[i-1]
	}
	return temp
}

func main() {
	a := []int{1, 2}
	a = insertMiddle(a, 100)
	fmt.Println(a)
	b := []int{1, 2, 3, 4}
	b = insertMiddle(b, 100)
	fmt.Println(b)

	// one loop version
	c := []int{1, 2}
	c = insertMiddleOneLoop(c, 100)
	fmt.Println(c)
	d := []int{1, 2, 3, 4}
	d = insertMiddleOneLoop(d, 100)
	fmt.Println(d)

	// idiomatic version
	e := []int{1, 2}
	e = insertMiddleGo(e, 100)
	fmt.Println(e)
	f := []int{1, 2, 3, 4}
	f = insertMiddleGo(f, 100)
	fmt.Println(f)

}
