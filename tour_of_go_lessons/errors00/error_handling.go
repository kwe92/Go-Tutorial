package main

import (
	"fmt"
	"time"
)

func main() {

	if err := new(); err != nil {
		fmt.Println(err)
	}
}

// the error built-in interface definition
type error interface {
	Error() string
}

func new() error {

	return &MyError{

		What: "This should throw an error.",
	}
}

type MyError struct {
	When time.Time

	What string
}

func (e *MyError) Error() string {

	e.When = time.Now()

	return fmt.Sprintf("at %v, %s", e.When, e.What)

}

// Error State

//   - Go has no traditional error handling
//   - errors are instead implemented with a built-in error interface
//   - the interface expects an implementation of an Error method that returns a string
//   - the process is similar to using the Stringer interface but:

//       - the Error method implements a pointer receiver instead of a value receiver

// Similarities to Stringer Interface

//   - a type is defined
//   - a method is implemented to represent that type as a string

// Process to Define Custom Errors

// 1. define a type
// 2. implement an Error method with a pointer receiver if returned from a function
//    a value receiver if instantiated directly that returns
//    a string representation of the type defined
