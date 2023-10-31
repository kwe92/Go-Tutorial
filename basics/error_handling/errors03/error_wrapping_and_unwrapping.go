package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func isEven(num int) (bool, error) {
	if odd := num%2 != 0; odd {
		return false, fmt.Errorf("expected an even number but received: %d", num)
	}
	return true, nil
}

func doubleEven(num int) (int, error) {
	_, err := isEven(num)
	if err != nil {
		return 0, fmt.Errorf("in doubleEven: %w", err)
	}
	return num * 2, nil
}

func main() {
	var evenNum = 42

	num, err := doubleEven(evenNum)

	fmt.Println(num)

	checkErrorAndUnwrap(err)

	var oddNum = 43

	num, err = doubleEven(oddNum)

	fmt.Println(num)

	checkErrorAndUnwrap(err)

}

func checkErrorAndUnwrap(err error) {
	if err != nil {
		fmt.Println(err)
		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {

			log.Fatalf("Wrapped Error Message: %s", wrappedErr.Error())

		}
		os.Exit(1)

	}
}

// Wrapping Errors

//   - prefixing errors with additional context e.g. the invoked function that caused the error

// fmt.Errorf | Error Wrapping

//   - formats a string based on format verbs / specifiers and returns an instantiated error
//   - if the format verb is %w then the error associated with the verb is wrapped

// Error Chain

//   - a chain of wrapped errors

// Unwrapping Errors

//   - use the errors.Unwrap method to return the result of calling the Unwrap method of the error passed in
//   - you can then check if this result is nil
//   - if the result is nil then there was no wrapped error
