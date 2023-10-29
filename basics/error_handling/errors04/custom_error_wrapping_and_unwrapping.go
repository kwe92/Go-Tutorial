package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

type InvalidStatus int

const (
	InvalidOdd  InvalidStatus = iota + 501
	InvalidEven InvalidStatus = iota + 502
)

type InvalidNumErr struct {
	InvalidStatus
	message string
	err     error
}

func (i InvalidNumErr) Error() string {
	return i.message
}

func (i InvalidNumErr) Unwrap() error {
	return i.err
}

func isEven(num int) (bool, error) {
	if odd := num%2 != 0; odd {
		return false, errors.New("not an even number.")
	}
	return true, nil
}

func doubleEven(num int) (int, error) {
	_, err := isEven(num)
	if err != nil {
		return 0, InvalidNumErr{
			InvalidStatus: InvalidOdd,
			message:       fmt.Sprintf("expected an even number and received %d", num),
			// initialize the received error into the custom errors error field
			err: err,
		}
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

// Wrapping Errors With Custom Errors

//   - the custom error struct should include a field that expects an error
//   - the error that you intend to wrap should be initialized to this field
//   - the custom error struct is also required to implement the Unwrap method
//     to be used with fmt.Unwrap, which expects the passed in error to have implemented this method
