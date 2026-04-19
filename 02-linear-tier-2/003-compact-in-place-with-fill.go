//go:build ignore
// +build ignore

package main

import "fmt"

/*
Name: Linear Compact In Place With Fill
Level: Easy
TC: O(n)  — requires TWO passes over the array (see below)
SC: O(1)
Desc: Compact non-zeros at the front in one pass, then fill the tail with

	zeros in a second pass. Unlike 002 (single pass, logical length),
	this returns the full slice fully cleaned up.

HOW this differs from 002:

	002 (no-fill) returns a logical length and leaves the tail as garbage.
	This version goes one step further: it explicitly zeroes out a[write:]
	so the caller gets back the full original slice — clean, with zeros
	pushed to the end. No logical length needed; the slice speaks for itself.

TWO-PASS approach:

	Pass 1 — same write-pointer technique as 002: scan with i, copy non-zeros
	          to the front, advance writer only on non-zero values.
	Pass 2 — fill a[writer:] with zeros to erase the leftover garbage tail.

	Both passes are O(n) in the worst case, but constants — overall TC stays O(n).

WHY still SC: O(1)?

	Same reason as 002: we only use two integer variables (writer, i).
	No new slice is allocated; the fill loop writes directly into the
	existing backing array. Memory used = constant, regardless of n.

Example trace for [1, 0, 2, 0, 3]:

	Pass 1 (write-pointer):
	  i=0: a[0]=1 != 0 → a[0]=1, writer=1
	  i=1: a[1]=0      → skip
	  i=2: a[2]=2 != 0 → a[1]=2, writer=2
	  i=3: a[3]=0      → skip
	  i=4: a[4]=3 != 0 → a[2]=3, writer=3
	After pass 1: [1, 2, 3, 0, 3]  ← tail is garbage (indices 3 and 4)

	Pass 2 (fill tail with zeros):
	  i=3: a[3]=0
	  i=4: a[4]=0
	Final: [1, 2, 3, 0, 0]  ← clean; caller can use the full slice
*/
func compactInPlaceWithFill(a []int) []int {
	writer := 0

	// Pass 1: move non-zeros to the front
	for i := range a {
		if a[i] != 0 {
			a[writer] = a[i] // overwrite gap (or same slot if no zeros seen yet)
			writer++
		}
		// zeros are skipped; writer stays behind, creating the gap
	}

	// Pass 2: zero out the garbage tail a[writer:]
	for i := writer; i < len(a); i++ {
		a[i] = 0
	}

	return a
}

func main() {
	fmt.Println(compactInPlaceWithFill([]int{1, 0, 2, 0, 3}))
}
