// package main contains example of comparable generic type usage
package main

import (
	"fmt"
)

var (
	intArr = []int{1, 2, 3, 4, 5, 6}
	strArr = []string{"first", "second", "apple", "banana"}
)

// ComparableSlice is an example of generic feature with new "comparable" type
type ComparableSlice[T comparable] []T

// AllEqual returns true when all slice elements are equal
// like [1, 1, 1] => true; [2, 1, 4] => false
func AllEqual[T comparable](s ComparableSlice[T]) bool {
	if len(s) == 0 {
		return true
	}

	last := s[0]
	for _, cur := range s[1:] {
		if cur != last {
			return false
		}
		last = cur
	}
	return true
}

// Contains returns true if slice arr contains element v
func Contains[T comparable](arr []T, v T) bool {
	for _, e := range arr {
		if e == v {
			return true
		}
	}
	return false
}

func main() {
	fmt.Printf("trying to find %v in given slice, result is: %v\n", 5,  Contains(intArr, 5))

	fmt.Printf("trying to find %v in given slice, result is: %v\n", "cat",  Contains(strArr, "cat"))

	fmt.Printf("trying to find %v in given slice, result is: %v\n", "apple",  Contains(strArr, "apple"))

	fmt.Printf("check all elemens of slice are equal: %v\n", AllEqual([]int{2,3,4,5}))

	fmt.Printf("check all elemens of slice are equal: %v\n", AllEqual([]string{"banana", "banana"}))

	//fmt.Printf("trying to find %v in given slice, result is: %v", 1,  Contains(strArr, 1))
}