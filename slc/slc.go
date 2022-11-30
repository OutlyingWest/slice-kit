// Here are useful techniques for working with slices that allow you to solve particular problems.
// Unlike other programming languages,
// Go does not flaunt an abundance of functions for working with slices.

package main

import (
	"fmt"
	"reflect"
)

type sliceTypes interface {
	int | int8 | int16 | int32 | float32 | ~string
}

// Delete the last slice element:
func DeleteLast[T sliceTypes](s []T) []T {
	if len(s) != 0 { // Protection from panic
		s = s[:len(s)-1]
	}
	return s
}

// Delete the first slice element:
func DeleteFirst[T sliceTypes](s []T) {
	if len(s) != 0 { // Protection from panic
		s = s[1:]
	}
	fmt.Println(s) // [2 3]
}

// Delete the slice element with index i.
// first and last element is not allowed.
// For delete first or last - use the DeleteFirst() or DeleteLast()
func DeleteInd[T sliceTypes](s []T, i int) []T {
	if len(s) != 0 && i < len(s)-1 { // Protection from panic
		s = append(s[:i], s[i+1:]...)
	}
	return s
}

// Return Values: Returns a slice of elements from start to end, inclusive.
func NewSlice(start, stop, step int) []int {
	if step <= 0 || stop < start {
		return []int{}
	}
	s := make([]int, 0, 1+(stop-start)/step)
	for start <= stop {
		s = append(s, start)
		start += step
	}
	return s
}

// Execute the reverse of slice elements
func ReverseSlice[T sliceTypes](s []T) []T {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
	return s
}

// Comparison of two slices:
func CompareSlices() {
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 4}
	s3 := []string{"1", "2", "3"}
	s4 := []int{1, 2, 3}

	fmt.Println(reflect.DeepEqual(s1, s2)) // false
	fmt.Println(reflect.DeepEqual(s1, s3)) // false
	fmt.Println(reflect.DeepEqual(s1, s4)) // true
}

// Delete duplications in slice
func Deduplicate[T sliceTypes](slice []T) []T {
	slcLen := len(slice)
	bufMap := make(map[T]int, slcLen)
	for slcInd, value := range slice {
		slcLenDec := slcLen - 1
		if _, ok := bufMap[value]; ok && slcInd != slcLenDec {
			slice = DeleteInd(slice, slcInd)
		} else if ok && slcInd == slcLenDec {
			slice = DeleteLast(slice)
		} else {
			bufMap[value] = slcInd
		}
	}
	return slice
}
