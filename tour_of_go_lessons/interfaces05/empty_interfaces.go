package main

import "fmt"

type I interface{}

func main() {
	var i I

	i = 42

	describe(i)

	i = "the LORD is my light and my salvation; whom shall I fear? the LORD is the strength of my life; of whom shall I be afraid?"

	describe(i)

}

func describe(i I) {

	fmt.Printf("\n(%v, %T)", i, i)

}

// The Empty Interface

//   - interface{} represents an empty interface type with possibly no methods
//     than can hold a value of any type which can be dynamically assigned
//   - the pointer pair (value, concrete_type) can be changed dynamically

// The any Type

//   - the value any has a static type of interface{}

// comma ok
