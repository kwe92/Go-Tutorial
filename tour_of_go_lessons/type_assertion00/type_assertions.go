package main

import "fmt"

type MyString string

func main() {

	var i0 interface{}

	var str00 MyString = "\nIn the beginning was the Word."

	i0 = str00

	i1 := i0.(MyString)

	// print the type of the underlying value
	fmt.Printf("\nType of i1: %T", i1)

	// will panic, the type must match the type of the underlying value | in this case MyString is the type of the underlying value
	// ! s := i0.(string)

	// fmt.Println(s)

	// comma ok statement: does not panic
	f, ok := i0.(float64)

	fmt.Printf("\nvalue: %v, isString: %v\n", f, ok)

}

// Type Conversions change, Type Assertions reveal

// Type Assertion | Checked at Runtime

//   - type assertions allow you to access the underlying concrete type of an interfaces value
//   - remember that an interface value is a pair of pointers (value, concrete_type) in the GO Run-time
//   - type is asserted and initializes declared variable with the interfaces value
//     if the assertion passes, if the assertion fails a panic is thrown

// To Panic or Not to Panic

//   - to not panic you can use the comma ok statement
//   - you should always validate your type assertions using the comma ok idiom
//   - if the assertion passes the interfaces value is returned along with a true value
//   - if the assertion fails the zero value of the asserted type is returned along with a false value
