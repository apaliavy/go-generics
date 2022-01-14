// package main contains a few examples of Sum functions using new generic union types
package main

import "fmt"

// Integer type combines all possible integet types
type Integer interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr
}

// SumInt64 doesn't bring anything new - just an example of sum function
func SumInt64(values ...int64) (result int64) {
	for _, v := range values {
		result += v
	}

	return result
}

// SumIntsOrFloats shows how to use types union in inline mode
func SumIntsOrFloats[V int64 | float64](values ...V) (result V) {
	for _, v := range values {
		result += v
	}
	return result
}

// SumIntegers does all the same like SumInt64 but works for our new generic Integer type
func SumIntegers[V Integer](values ...V) (result V) {
	for _, v := range values {
		result += v
	}
	return result
}

func main()  {
	fmt.Printf("calling SumInt64: %v\n", SumInt64(int64(10), int64(20), int64(30)))
	fmt.Printf("calling SumIntsOrFloats: %v\n", SumIntsOrFloats(int64(10), int64(20), int64(30)))
	fmt.Printf("calling SumIntsOrFloats: %v\n", SumIntsOrFloats(float64(10.10), float64(20.04), float64(30.32)))

	//fmt.Printf("calling SumIntsOrFloats: %v\n", SumIntsOrFloats(int64(10), float64(20.04), float64(30.32)))

	fmt.Printf("calling SumIntegers: %v\n", SumIntegers(20, 10, 15))

	//fmt.Printf("calling SumIntegers: %v\n", SumIntegers(20, 10, int64(11), int32(15)))
}