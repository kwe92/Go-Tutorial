package main

// TODO: comment code

import (
	"errors"
	"fmt"
	"strconv"
)

func doubleEven(num int) (int, error) {
	if odd := num%2 != 0; odd {
		return 0, errors.New("in doubleEven: Expected even number received odd.")
	}
	return num * 2, nil
}

func parseInt(num string) (int, error) {
	parsedNum, err := strconv.Atoi(num)

	if err != nil {
		return 0, fmt.Errorf("in parseInt: %w", err)
	}

	return parsedNum, nil
}

func div(numerator, denominator int) (int, error) {
	if denominator == 0 {
		return 0, errors.New("in div: divide by zero error.")
	}
	return numerator / denominator, nil
}

func DoThings(val1 int, val2 string) (_ int, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("in DoThings: %w", err)
		}
	}()

	val3, err := doubleEven(val1)

	if err != nil {
		return 0, err
	}

	val4, err := parseInt(val2)

	if err != nil {
		return 0, err
	}

	finalValue, err := div(val3, val4)

	return finalValue, err
}

func main() {

	valuesMap := map[int]string{
		12:  "4",
		64:  "32",
		300: "0",
		55:  "5",
		56:  "not a number",
	}

	for num, stringNum := range valuesMap {
		result, err := DoThings(num, stringNum)

		fmt.Printf("Result: %d | Error: %v\n\n", result, err)
	}

}

// Wrapping Multiple Errors With Same Message

//   - if you have a multiple of errors that could be wrapped with the same message
//     you can use defer to wrap the message keeping code DRY
