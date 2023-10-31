package main

import (
	"errors"
	"fmt"
	"os"
)

type FileNotExistErr struct {
	message string
	err     error
}

func (f FileNotExistErr) Error() string {
	return f.message
}

func (f FileNotExistErr) Unwrap() error {
	return f.err
}

// fileAsString: returns the entire contents of a file as a string
func fileAsString(filepath string) (string, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return "", FileNotExistErr{
			message: fmt.Sprintf("in fileAsString: %v", err),
			err:     err,
		}
	}

	stringData := string(data)

	return stringData, nil

}

func main() {

	var filePath00 string = "basics/error_handling/errors05/some_text_file.txt"
	var filePath01 string = "i_dont_exist.txt"

	fileStringRep, err := fileAsString(filePath00)

	checkError(err)

	fmt.Println(fileStringRep)

	fileStringRep, err = fileAsString(filePath01)

	if err != nil {
		fmt.Println(errors.Is(err, os.ErrNotExist))
		if notExistErr := errors.Is(err, os.ErrNotExist); notExistErr {
			fmt.Println("the file or directory path does not exist.")
		}
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// Problem of Wrapping Errors

//   - you can not use the equality operator == to check for wrapped sentinal errors
//   - you can not use type assertion or type switch to match custom wrapped errors

// errors.Is | With Sentinal Errors

//   - takes two arguments the error received and the target error
//   - errors.Is then repeatedly calls the Unwrap method until there is a match
//     or there are no more errors to unwrap
//   - Typically the error provided is a sentinal error

// errors.Is | With Noncomparable type

//   - if the type of your custom error is noncomparable then you must implement the Is method
//     on your custom error
