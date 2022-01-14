// package main contains example of any type usage
package main

import "fmt"

// Filter accepts slice of any type and returns result of filtration.
// Use have to provide your own filter function f
func Filter[T any](a []T, f func(T) bool) []T {
	var n []T
	for _, e := range a {
		if f(e) {
			n = append(n, e)
		}
	}
	return n
}

func main() {
	vi := []int{1, 2, 3, 4, 5, 6}

	vi = Filter(vi, func(v int) bool {
		return v > 3
	})

	fmt.Println(vi)
}