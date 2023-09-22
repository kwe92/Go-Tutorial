package main

import "fmt"

func main() {

	i0 := interface{}("\nIn the beginning was the Word.")

	i1 := any("In him was life; and the life was the light of men.")

	i1 = 42

	s := i0.(string)

	fmt.Println(s)

	// comma ok statement: does not panic
	f, ok := i0.(float64)

	fmt.Printf("\nvalue: %v, isString: %v\n", f, ok)

	fmt.Println(i1.(int))

	// panics!!!
	f = i0.(float64)

	fmt.Println(f)

}

// Type Assertion

//   - type assertions allow you to access the underlying concrete type of an interfaces value
//   - remeber that an interface value comes as a pair of pointers (value, concrete_type)
//   - type is asserted and initializes declared variable with the interfaces value
//     if the assertion passes, if the assertion fails a panic is thrown

// To Panic or Not to Panic

//   - to not panic you can use the comma ok statement
//   - if the assertion passes the interfaces value is returned along with a true value
//   - if the assertion fails the zero value of the asserted type is returned along with a false value
