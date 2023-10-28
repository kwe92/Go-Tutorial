package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

// doubleEven: doubles a number if it is even
func doubleEven(num int) (int, error) {
	if odd := num%2 != 0; odd {
		return 0, errors.New("expected even number, received odd.")
	}
	doubledNum := num * 2
	return doubledNum, nil
}

// trippleOdd: doubles a number if it is even
func trippleOdd(num int) (int, error) {
	if even := num%2 == 0; even {
		return 0, fmt.Errorf("expected odd number, received: %d.", num)
	}
	doubledNum := num * 3
	return doubledNum, nil
}

func main() {
	evenNum := 42

	num, err := doubleEven(evenNum)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(num)

	oddNum := 43

	num, err = doubleEven(oddNum)

	if err != nil {
		log.Fatalf(err.Error())
	}

	num, err = trippleOdd(evenNum)

	if err != nil {
		log.Fatalf(err.Error())

	}

}

// Error Handling in GO

//   - GO does not have try catch or finally keywords for the caller to handle errors

//   - errors are handled explicitly by the caller of a function
//     by checking if the error interface is returned nil or not

// Stopping Execution When Encountering an Error

//   - there are a few ways to stop the execution of a program if an error is encountered

// Stopping Execution with fmt.Printf and os.Exit(1)

//   - print the error to the console and call os.Exit(1) to exit the
//     current program with an non-zero error code to indicate failure

// Stopping Execution with log.Fatalf(error.Error())

//   - passing the error string to log.Fatal or log.Fatalf for string formating
//   - has the same effect as using fmt.Println or fmt.Printf with os.Exit(1) with additional capabilities
//   - prefixes the error with a timestamp

// Why Check for and return nil?

//   - the return type of an error is an interface
//   - the zero value for any interface type is nil

// Returning Errors

//   - there are a few built-in ways to return errors
//   - you can also return your own error that implements the error interface

// errors.New

//   - takes a string without interpolation / formating and returns a new error interface

// fmt.Errorf

//   - takes a string with interpolation / formating and returns a new error interface
//   - allows you to use the formatting verbs / specifiers
