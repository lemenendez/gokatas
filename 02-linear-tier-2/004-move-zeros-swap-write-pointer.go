//go:build ignore
// +build ignore

package main

import "fmt"

/*
Name: Linear Move Zeros Right (Swap + Write Pointer)
Level: Easy
TC: O(n)
SC: O(1) extra
Desc: Move all zeros to the right in-place by swapping non-zeros forward.

Interview signal:
	This is a very common coding interview pattern.
	Why interviewers like it:
	- tests in-place array manipulation (no extra array allowed)
	- tests pointer/index invariants under mutation
	- tests ability to preserve relative order while transforming data
	- invites follow-ups (count swaps, return logical length, mirror direction)

Typical prompt wording:
	"Move all zeros to the end while keeping non-zero order, in-place."
	"Can you do it in O(n) time and O(1) extra space?"

Goal:
	Keep non-zero values in their original relative order (stable behavior),
	and push all zeros to the end of the same slice.

Why this is O(1) extra:
	You only use a couple of integer indexes (for example, `write` and `i`).
	No second slice is allocated.

Approach hint (single pass + swaps):
	- `write` points to the next slot where a non-zero should land.
	- `i` scans from left to right.
	- When a[i] is non-zero, swap a[write] and a[i], then increment write.
	- When a[i] is zero, skip it (only i moves).

Invariant idea to keep in mind while coding:
	- a[:write] contains only non-zeros, in original order.
	- a[write:i+1] is the active region still being processed.

Interview tip:
	When explaining your solution out loud, state the invariant first,
	then walk one short example. That usually scores better than only code.
*/

// TODO:
// 1) Initialize write = 0.
// 2) Loop i from 0 to len(a)-1.
// 3) If a[i] != 0, swap a[write] with a[i], then write++.
// 4) Return the mutated slice.
// 5) Mentally verify with: [1,0,3,0,4] -> [1,3,4,0,0].
func moveZerosSwapWritePointer(a []int) []int {
	writer := 0
	for i := 0; i < len(a); i++ {
		if a[i] != 0 {
			a[writer], a[i] = a[i], a[writer]
			writer++
		}
	}
	return a
}

func main() {
	fmt.Println(moveZerosSwapWritePointer([]int{1, 0, 2, 0, 3}))
}
