package main

import "fmt"

// GetIndex compares x to all elements in slice, returns index of x if exists or -1
func GetIndex[T comparable](slice []T, x T) int {

	for index, value := range slice {
		if value == x {
			return index
		}
	}
	return -1
}

func main() {
	// both strings and number implement the comparable interfac

	intSlice := []int{0, 3, 6, 9, 12, 15, 18, 21, 24}

	strSlice := []string{"welcome", "the", "void"}

	fmt.Println(GetIndex(intSlice, 9))

	fmt.Println(GetIndex(strSlice, "Universe"))

}

// Go Type Parameters

//   - allows Go functions to work with multiple types
//   - type parameters appear between the function name and function parameters
//     placed between brackets
//   - parameterized functions with constraints

// comparable built-in constraint

//   - an interface that is implemented by all comparable types such as:
//       (booleans, strings, numbers, collections of comparable types, struct with comparable type fields)
//   - can only be used as a type parameter
//   - allows the use of the equality and inequality operators == && !=
//     for comparason of parameterized types
//   - incorrect types are caught at compile-time
