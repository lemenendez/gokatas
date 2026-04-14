//go:build ignore
// +build ignore

package main

import (
	"errors"
	"fmt"
)

var ErrOutOfBounds = errors.New("out of bounds")

// Array is a learning-oriented linear data structure.
//
// Intention: implement operations manually (insert/delete/append/reverse/etc.)
// without using built-in helpers like append or copy.
type Array struct {
	// len tracks the logical number of valid elements.
	// It can be smaller than cap(data) when preallocating space.
	len int
	// data stores the backing memory used by this custom structure.
	data []int
}

// NewArray creates an empty Array instance.
// Time: O(1)
// Space: O(1)
func NewArray() *Array {
	return &Array{
		data: make([]int, 0),
	}
}

// Len returns the logical number of stored elements.
// Time: O(1)
// Space: O(1)
func (a *Array) Len() int {
	return a.len
}

// Get returns the value at index.
// Time: O(1)
// Space: O(1)
func (a *Array) Get(index int) (int, error) {
	if index < 0 || index >= a.len {
		return 0, ErrOutOfBounds
	}
	return a.data[index], nil
}

// Set updates the value at index.
// Time: O(1)
// Space: O(1)
func (a *Array) Set(index int, value int) error {
	if index < 0 || index >= a.len {
		return ErrOutOfBounds
	}
	a.data[index] = value
	return nil
}

// Append adds a value at the end without using append/copy.
// Time: O(n), because n elements are copied to new storage.
// Space: O(n) extra for the temporary array.
func (a *Array) Append(value int) {
	l := a.len
	temp := make([]int, l+1)
	for i := 0; i < l; i++ {
		temp[i] = a.data[i]
	}
	temp[l] = value
	a.len++
	a.data = temp
}

// InsertAt inserts value at index without using append/copy.
// Note: index == len is allowed and behaves like append.
// Time: O(n), because values are copied around the insertion point.
// Space: O(n) extra for the temporary array.
func (a *Array) InsertAt(index int, value int) error {
	if index < 0 || index > a.len {
		return ErrOutOfBounds
	}
	l := a.len
	temp := make([]int, l+1)

	// Copy left side [0, index).
	for i := 0; i < index; i++ {
		temp[i] = a.data[i]
	}

	// Place new value.
	temp[index] = value

	// Copy right side [index, l) shifted by +1.
	for i := index + 1; i < l+1; i++ {
		temp[i] = a.data[i-1]
	}

	a.data = temp
	a.len++
	return nil
}

// DeleteAt removes the value at index and returns it.
// Time: O(n), because values after index are shifted left.
// Space: O(n) extra for the temporary array.
func (a *Array) DeleteAt(index int) (int, error) {
	if index < 0 || index >= a.len {
		return 0, ErrOutOfBounds
	}

	temp := make([]int, a.len-1)

	// Copy left side [0, index).
	for i := 0; i < index; i++ {
		temp[i] = a.data[i]
	}

	value := a.data[index]

	// Copy right side (index, len) shifted by -1.
	for i := index + 1; i < a.len; i++ {
		temp[i-1] = a.data[i]
	}

	a.len--
	a.data = temp
	return value, nil
}

// ReverseInPlace reverses the current logical elements.
// Time: O(n), about n/2 swaps.
// Space: O(1) extra.
func (a *Array) ReverseInPlace() {
	for i, j := 0, a.len-1; i < a.len/2; i, j = i+1, j-1 {
		a.data[i], a.data[j] = a.data[j], a.data[i]
	}
}

// Count returns how many times value appears.
// Time: O(n)
// Space: O(1)
func (a *Array) Count(value int) int {
	c := 0
	for i := 0; i < a.len; i++ {
		if a.data[i] == value {
			c++
		}
	}
	return c
}

// Search returns the first index of value, or -1 if not found.
// Time: O(n) in the worst case.
// Space: O(1)
func (a *Array) Search(value int) int {
	for i := 0; i < a.len; i++ {
		if a.data[i] == value {
			return i
		}
	}
	return -1
}

func main() {
	a := NewArray()

	a.Append(10)
	a.Append(20)
	a.Append(30)
	fmt.Println("after append:", a.data, "len:", a.Len())

	err := a.InsertAt(1, 15)
	fmt.Println("insert err:", err, "data:", a.data, "len:", a.Len())

	v, err := a.Get(2)
	fmt.Println("get index 2:", v, "err:", err)

	err = a.Set(2, 99)
	fmt.Println("set err:", err, "data:", a.data)

	deleted, err := a.DeleteAt(1)
	fmt.Println("deleted:", deleted, "err:", err, "data:", a.data, "len:", a.Len())

	fmt.Println("count 99:", a.Count(99))
	fmt.Println("search 30:", a.Search(30))
	fmt.Println("search 404:", a.Search(404))

	a.ReverseInPlace()
	fmt.Println("after reverse:", a.data)

	_, err = a.Get(999)
	fmt.Println("out of bounds get err:", err)
}
