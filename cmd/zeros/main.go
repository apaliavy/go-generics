// package main contains example of work with zero values with generic T type
package main

import "fmt"

// ReturnZero demonstrates how nil-values work with generic methods and types
func ReturnZero[T any](s ...T) T {
	var zero T
	return zero
}

func main() {
	fmt.Println(ReturnZero(5))
	// prints "0"
	fmt.Println(ReturnZero("string"))
	// prints ""
	fmt.Println(ReturnZero(true))
	// prints "false"
}
