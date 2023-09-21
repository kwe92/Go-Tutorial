package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I

	describe(i)

	// Will throw a run-time error
	i.M()
}

func describe(i I) {

	fmt.Printf("\n%v, %T\n", i, i)

}

// Nil Interface Value

//   - interfaces declared without initialization hold neither value nor concrete type.
//   - Calling methods on nil interfaces raises run-time errors
//     as value and concrete type are nil pointing to neither data nor concrete type in memory
