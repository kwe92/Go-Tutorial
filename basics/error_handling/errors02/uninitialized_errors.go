package main

import "fmt"

type InvalidStatus int

const (
	InvalidNumberOdd InvalidStatus = iota + 501
	InvalidNumberEven
)

type InvalidNumberErr struct {
	InvalidStatus
	message string
}

func (i InvalidNumberErr) Error() string {
	return i.message
}

func generateErrorWrong(flag bool) error {

	// InvalidNumberErr is a struct not an interface and it's fields are initialized to their zero-value
	var invalidNumErr InvalidNumberErr

	if flag {
		invalidNumErr = InvalidNumberErr{
			message: "Invalid number.",
		}

		return invalidNumErr
	}

	return invalidNumErr

}

func generateErrorRight(flag bool) error {

	// error is an interface and initialized with the zero-value of nil which also makes its concrete type nil
	var invalidNumErr error

	if flag {
		invalidNumErr = InvalidNumberErr{
			message: "Invalid number.",
		}

		return invalidNumErr
	}

	return invalidNumErr

}

// generateErrorRight2: is the cleanest and shortest way to return an uninitialized error

func generateErrorRight2(flag bool) error {

	if flag {
		return InvalidNumberErr{
			message: "Invalid number.",
		}

	}

	return nil

}

func main() {

	var t bool = true

	// zero-value of bool is false
	var f bool

	err := generateErrorWrong(t)

	fmt.Printf("\nexpected %t received %t", t, err != nil)

	err = generateErrorWrong(f)

	fmt.Printf("\nexpected %t received %t", f, err != nil)

	err = generateErrorRight(t)

	fmt.Printf("\nexpected %t received %t", t, err != nil)

	err = generateErrorRight(f)

	fmt.Printf("\nexpected %t received %t", f, err != nil)

	err = generateErrorRight2(t)

	fmt.Printf("\nexpected %t received %t", t, err != nil)

	err = generateErrorRight2(f)

	fmt.Printf("\nexpected %t received %t\n", f, err != nil)
}

// Uninitialized Instances of Custom Errors

//   - custom errors are user defined types not interfaces
//   - if a custom error is declared but not initialized then
//     the fields of the custom error are initialized to their zero values

// Avoiding Uninitialized Instances of Custom Errors

//   - always declare an error variable as an error interface
//   - instantiate any implementation of the error interface as the variables value
//   - return nil if an error was declared as a variable but there was not an instantiation of an error implementation
