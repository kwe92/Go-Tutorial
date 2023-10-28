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

	checkErr(err)

	var oddNum = 43

	num, err = doubleEven(oddNum)

	fmt.Println(num)

	checkErrorAndUnwrap(err)

	// checkErr(err)

}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func checkErrorAndUnwrap(err error) {
	if err != nil {
		fmt.Println(err)
		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
			fmt.Println("Wrapped Error Message:", wrappedErr)
			os.Exit(1)
		}
		os.Exit(1)

	}
}

// Wrapping Errors

//   - adding additional context to errors

// fmt.Errorf | Error Wrapping

//   - can format a string based on format verbs
//   - if the format verb is %w then the error associated with the verb is wrapped

// Error Chain

//   - a chain of wrapped errors
