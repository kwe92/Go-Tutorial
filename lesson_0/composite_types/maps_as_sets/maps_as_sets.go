package main

import "fmt"

// a set of integers.
type IntSet map[int]bool

// Values: returns a Slice of all elements within the set.
func (s *IntSet) Values() []int {
	arr0 := make([]int, 0)
	for k := range *s {
		arr0 = append(arr0, k)
	}
	return arr0
}

func main() {
	// Slice of integers
	intArr := []int{2, 2, 4, 4, 5, 6, 7, 8, 9, 9, 11, 12, 13, 13}

	intSet := IntSet{}

	for _, ele := range intArr {
		intSet[ele] = true
	}

	var intSetValues []int

	intSetValues = intSet.Values()

	fmt.Println(len(intArr), len(intSet))

	fmt.Println(intArr)

	fmt.Println(intSetValues)

}

// Sets

//   - A collection of unique elements

// Sets in GO

//   - the set data structure does not exist in GO
//   - maps and structs can be used as sets
//   - maps are prefered over structs
//   - set operations must be implemented by the developer or attained from a third party package

// maps as Sets

//   - to use a map as a set you can assign the maps keys comparable types and the values of type boolean
//   - keys in a map must be unique

// map for-range statement

//   - the for range statement of a map can potentially reorder the elements
//   - hash map implementations are typically unordered key-value pairs in most programming languages
