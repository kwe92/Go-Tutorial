package main

import "fmt"

func whatTheType(i interface{}) {

	switch v := i.(type) {

	case int:
		fmt.Printf("\nvalue: %v, type: %T", v, v)

	case string:
		fmt.Printf("\nvalue: %q | value is %v bytes long", v, len(v))

	default:
		fmt.Printf("\nI am not sure what the type of %v is", v)

	}

}

func forEach[T any](arr []T, fn func(T)) {

	for _, ele := range arr {
		fn(ele)
	}
}

var s string

var integer int

var float float64

func main() {

	s = "There was a man sent from God, whose name was john."

	integer = 42

	float = 3.14

	arr0 := []interface{}{s, integer, float}

	forEach(arr0, whatTheType)

	fmt.Println()
}

// Type Switch Statements

//   - used when an interface can be one of multiple types
//   - allows you to evaluate several type assertions in a series of case statements
//   - the cases are types instead of values
//   - the case compares the type to the interface values underlying concrete type
//   - the variable being checked is usually assigned to a variable of the same name, but doesnt have to be
//     this is a situation where shadowing a variable name is okay

// Type Swtich Declaration

//   - the declaration of a type switch is the same as a type assertion with one difference:

//       - the type passed as an assertion is the `type` keyword e.g. switch v := i.(type)
