package main

import "fmt"

func main() {

	var i I

	var t *T

	i = t

	describe(i)

	i.M()

	i = &T{S: "Have I not commmanded thee? Be strong and of good courage, be not afraid, neither be thou dismayed: for the LORD thy God is with thee whithersoever thou goest."}

	describe(i)

	i.M()

}

func describe(i I) {
	fmt.Printf("(%v, %T)", i, i)
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("\n<nil>")
		return
	}
	fmt.Println("\n\n", t.S)
}

// Interface Values With nil Underlying Values

//   - If an interface is passed a concrete value that was declared
//     but not initialized with a value then the associated methods are passed a nil receiver
//   - In GO a null pointer exception is not thrown developers have the option
//     to handle nil receivers gracefully within the method implementation with control flow statements
//   - Interface values that hold a concrete nil value are not null and return their type
