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

	twoBase, threeBase := makeMulti(2), makeMulti(3)

	for i := 0; i < 5; i++ {
		fmt.Printf("\ntwoBase: %d | threeBase:  %d", twoBase(i), threeBase(i))

	}

	// NOTE: the sum parameter is by value not by reference so a copy is created
	fmt.Println("\nsum value: ", sum)

}

func adder(sum int) IntCallback {
	return func(x int) int {
		sum += x
		return sum
	}
}

func makeMulti(base int) IntCallback {
	return func(factor int) int {
		return base * factor
	}
}

// Function closures

//   - Go functions may be closures
//   - closures are function values that reference mutable variables bound to them outside of their own scope
//   - the state of parent function parameters are saved (bound) as constants for the lifetime of the closure
//   - e.g. the adder function returns a closure in which each closure is bound to its own sum variable.

// Binding | Lexical Scope

// Shadowed Variables

//   - variables that have the same name in diffrent scopes / blocks
