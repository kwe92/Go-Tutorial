package main

import "fmt"

const (
	InvalidNumberOdd  Invalid = iota + 501
	InvalidNumberEven Invalid = iota + 502
)

type Invalid int

type InvalidNumErr struct {
	Invalid
	message string
}

func (i InvalidNumErr) Error() string {
	return i.message
}

func generateErrorWrong(flag bool) error {

	// initialized with the zero value of empty string
	var invalidNumErr InvalidNumErr

	if flag {
		invalidNumErr = InvalidNumErr{
			message: "Invalid number.",
		}

		return invalidNumErr
	}

	return invalidNumErr

}

func generateErrorRight(flag bool) error {

	// initialized with the zero value of nil
	var invalidNumErr error

	if flag {
		invalidNumErr = InvalidNumErr{
			message: "Invalid number.",
		}

		return invalidNumErr
	}

	return invalidNumErr

}

func generateErrorRight2(flag bool) error {

	// initialized with the zero value of nil

	if flag {
		return InvalidNumErr{
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

	fmt.Printf("\nexpected %v received %v", t, err != nil)

	err = generateErrorWrong(f)

	fmt.Printf("\nexpected %v received %v", f, err != nil)

	err = generateErrorRight(t)

	fmt.Printf("\nexpected %v received %v", t, err != nil)

	err = generateErrorRight(f)

	fmt.Printf("\nexpected %v received %v", f, err != nil)

	err = generateErrorRight2(t)

	fmt.Printf("\nexpected %v received %v", t, err != nil)

	err = generateErrorRight2(f)

	fmt.Printf("\nexpected %v received %v\n", f, err != nil)
}

// Uninitialized Instances of Custom Errors

//   - custom errors are types not interfaces
//   - if a custom error is declared but not initialized then the zero value
//     of its underlying type will be initialized

// Avoiding Uninitialized Instances of Custom Errors

//   - always declare an error variable as an error interface and instantiate any implementation of the error interface as the variables value
//   - return nil if an error was declared as a variable but there was not an instantiation of an error implementation
