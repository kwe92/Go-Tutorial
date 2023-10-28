package main

import (
	"fmt"
	"log"
)

type Invalid int

const (
	InvalidNumberOdd  Invalid = iota + 501
	InvalidNumberEven Invalid = iota + 501
)

type InvalidNumberErr struct {
	Invalid
	message string
}

func (i InvalidNumberErr) Error() string {
	return i.message
}

// doubleEven: doubles a number if it is even
func doubleEven(num int) (int, error) {
	if odd := num%2 != 0; odd {
		return 0, InvalidNumberErr{
			Invalid: InvalidNumberOdd,
			message: "expected even number, received odd.",
		}
	}
	doubledNum := num * 2
	return doubledNum, nil
}

// trippleOdd: doubles a number if it is even
func trippleOdd(num int) (int, error) {
	if even := num%2 == 0; even {
		return 0, InvalidNumberErr{
			Invalid: InvalidNumberEven,
			message: fmt.Sprintf("expected odd number, received: %d.", num),
		}

	}
	doubledNum := num * 3
	return doubledNum, nil
}

func main() {
	evenNumber := 42

	num, err := doubleEven(evenNumber)

	checkErr(err)

	fmt.Println("Number: ", num)

	oddNumber := 43

	num, err = doubleEven(oddNumber)

	checkErr(err)

	fmt.Println("Number: ", num)

}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Custom Error

//   - error is an interface that can be implemented with a user defined-type that implements the Error method
//   - the names of custom errors should be suffixed with `Err`
//   - interfaces are structurally typed and defined-types that
//     implement an interface can have additional information

// Returning Custom Errors From Functions

//   - you should always return the error interface as the error result
//     instead of the custom defined error as the return type
