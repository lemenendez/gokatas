//go:build ignore
// +build ignore

package main

import "fmt"

/*
Name: Linear Filter Pattern (Remove matching elements)
Level: Easy
TC: O(n)
SC: O(1) extra
Desc: Generic in-place filter using swap + write-pointer to remove elements
      matching a predicate. Generalizes the move-zeros pattern to any predicate.

Goal:
	Keep only elements where the predicate returns false.
	Remove (push to the end) all elements where the predicate returns true.
	Preserve relative order of kept elements.

Why this matters:
	This shows how the swap + write-pointer technique is not just for zeros.
	Once you master the pattern, you can filter anything: negatives, odds,
	values > threshold, etc. In production code, this is how in-place filters work.

Predicate semantics:
	The predicate f(v) returns true if the element should be REMOVED.
	If f(v) == false, the element is KEPT and moved to the front.

How it differs from the move-zeros specific solutions:
	Instead of hardcoding "if a[i] != 0", we pass a function that decides
	what to remove. This is reusable and tests abstraction thinking.

Interview angle:
	"Can you generalize your move-zeros solution to work with any predicate?"
	This question often follows the simpler version in interviews.
*/

// removeWithFilter removes elements where f(v) returns true, keeps others.
// Returns a slice of the kept elements in original order, in-place.
func removeWithFilter(a []int, f func(i int) bool) []int {
	writer := 0
	for i := 0; i < len(a); i++ {
		if !f(a[i]) {
			a[writer], a[i] = a[i], a[writer]
			writer++
		}
	}
	return a[0:writer]
}

func isZero(v int) bool { return v == 0 }
func isNeg(v int) bool  { return v < 0 }
func isOdd(v int) bool  { return v%2 != 0 }

func main() {
	// Example 1: remove zeros (keep non-zeros)
	fmt.Println(removeWithFilter([]int{1, 3, 0, 4}, isZero))

	// Example 2: remove negative numbers
	fmt.Println(removeWithFilter([]int{1, 3, -1, 4}, isNeg))

	// Example 3: remove odd numbers (keep only even)
	fmt.Println(removeWithFilter([]int{1, 3, -1, 4}, isOdd))
}
