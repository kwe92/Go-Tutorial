package main

import (
	"fmt"
)

type IntCallback func(int) int

func main() {

	var sum int = 0

	// declare and initialize two separate adder functions
	pos, neg := adder(sum), adder(sum)

	for i := 1; i < 11; i++ {
		fmt.Printf("\nPos: %d | Neg %d", pos(i), neg(-2*i))
	}

	// NOTE: the sum parameter is by value not by reference so a copy is created

	fmt.Println("\nsum value: ", sum)
}

func adder(sum int) IntCallback {
	// var sum int = 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// Function closures

//   - Go functions may be closures
//   - closures are function values that reference mutable variables bound to them outside of their own scope
//   - e.g. the adder function returns a closure in which each closure is bound to its own sum variable.

// Binding | Lexical Scope
