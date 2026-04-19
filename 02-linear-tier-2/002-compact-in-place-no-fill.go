//go:build ignore
// +build ignore

package main

import "fmt"

/*
Name: Linear Compact In Place Without Fill
Level: Easy
TC: O(n)
SC: O(1) extra
Desc: Compact non-zeros at the front and return the logical length.

WHY "O(1) extra"?
  "Extra" (also called "auxiliary") space means memory beyond the input itself.
  Here we only declare two plain integer variables — `write` and `i` — regardless
  of how large the input is. Whether the slice has 10 or 10 000 000 elements,
  we still use exactly those two variables. That constant overhead is O(1).

  Contrast this with 001-collect-non-zeros-new-array.go, which calls
  make([]int, 0, len(a)) — allocating a second slice that grows with the input,
  making its extra space O(n). The benchmark in bench/bench.go makes this
  concrete: the O(n) variant allocates ~8 bytes per element on the heap,
  while this one shows 0 KB of extra heap allocation.

HOW the write pointer works:
  `write` tracks the next position where a non-zero value should land.
  `i` scans forward independently. When a[i] != 0, we copy it to a[write]
  and advance write. When a[i] == 0, we skip — write stays put, creating
  a "gap" that will be overwritten by the next non-zero found by i.

  Example trace for [1, 0, 3, 0, 4]:
    i=0: a[0]=1 != 0 → a[0]=1, write=1
    i=1: a[1]=0      → skip
    i=2: a[2]=3 != 0 → a[1]=3, write=2   ← overwrites the gap at index 1
    i=3: a[3]=0      → skip
    i=4: a[4]=4 != 0 → a[2]=4, write=3
  Result slice: a[:3] = [1, 3, 4]

NOTE: the tail of the original slice (a[write:]) is NOT zeroed out.
  That is intentional — the caller uses the returned logical length `write`
  to know the valid region. See 003 for the variant that fills the tail.
*/
func compactInPlaceNoFill(a []int) int {
	// write = next index available for a non-zero value
	write := 0

	for i := range a {
		if a[i] != 0 {
			a[write] = a[i] // overwrite the gap (or same position if no zeros yet)
			write++
		}
		// zero values are simply skipped; write does not advance
	}
	// everything in a[write:] is leftover garbage — don't use it
	return write
}

func main() {
	a := []int{1, 0, 3, 0, 4}
	n := compactInPlaceNoFill(a)
	fmt.Println(a[:n]) // [1 3 4]

	b := []int{0, 0, 0}
	n = compactInPlaceNoFill(b)
	fmt.Println(b[:n]) // []

	c := []int{1, 2, 3}
	n = compactInPlaceNoFill(c)
	fmt.Println(c[:n]) // [1 2 3]

	d := []int{}
	n = compactInPlaceNoFill(d)
	fmt.Println(d[:n]) // []
}
