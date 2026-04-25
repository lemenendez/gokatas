package main

import "fmt"

/*
Name: Remove Duplicates (Sorted Invariant + Dictionary)
Level: Easy
TC: O(n)
SC: O(1) extra for invariant version, O(n) extra for dictionary version
Desc: Two in-place dedup strategies:
      1) Sorted invariant (adjacent compare + write pointer)
      2) Dictionary/set-assisted dedup for unsorted input

Goal:
	Compact unique values to the front and return the compacted prefix.

Why this matters:
	This kata contrasts two important approaches.
	- Sorted invariant teaches pointer invariants with minimal space.
	- Dictionary version handles general input while preserving first-seen order.
*/

// removeDupsInvariant removes duplicates in-place from a sorted slice.
// Precondition: equal values are adjacent.
func removeDupsInvariant(a []int) []int {
	if len(a) == 0 {
		return a
	}
	// k writer, boundary of the clean prefix
	k := 1
	for i := 1; i < len(a); i++ {
		if a[i] != a[k-1] {
			a[k] = a[i]
			k++
		}
	}
	return a[0:k]
}

// removeDupsWithDict removes duplicates in-place from any input order.
// It preserves the first-seen order of unique values.
func removeDupsWithDict(a []int) []int {
	if len(a) == 0 {
		return a
	}
	seen := make(map[int]bool)
	k := 0
	for i := 0; i < len(a); i++ {
		if !seen[a[i]] {
			a[k] = a[i]
			seen[a[i]] = true
			k++
		}
	}
	return a[0:k]
}

func main() {
	/*
		fmt.Println(removeDupsInvariant([]int{}))
		fmt.Println(removeDupsInvariant([]int{1, 1}))
		fmt.Println(removeDupsInvariant([]int{1, 2, 2}))
		fmt.Println(removeDupsInvariant([]int{1, 1, 2, 2}))
		fmt.Println(removeDupsInvariant([]int{1, 1, 1}))
		fmt.Println(removeDupsInvariant([]int{1, 2, 3, 4, 5, 5, 6, 6}))
		fmt.Println(removeDupsInvariant([]int{5, 5, 6, 6}))
	*/

	fmt.Println(removeDupsWithDict([]int{}))
	fmt.Println(removeDupsWithDict([]int{1, 1}))
	fmt.Println(removeDupsWithDict([]int{1, 2, 2}))
	fmt.Println(removeDupsWithDict([]int{1, 1, 2, 2}))
	fmt.Println(removeDupsWithDict([]int{1, 1, 1}))
	fmt.Println(removeDupsWithDict([]int{1, 2, 3, 4, 5, 5, 6, 6}))
	fmt.Println(removeDupsWithDict([]int{5, 5, 6, 6}))
}
