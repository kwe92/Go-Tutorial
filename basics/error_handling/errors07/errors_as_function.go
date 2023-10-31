package main

import (
	"errors"
	"fmt"
)

type InvalidStatus int

const (
	InvalidNumOdd InvalidStatus = iota + 501
	InvalidNumEven
)

type InvalidNumErr struct {
	InvalidStatus
	message string
}

func (i InvalidNumErr) Error() string {
	return i.message
}

func isEven(num int) (bool, error) {
	if odd := num%2 != 0; odd {
		return false, InvalidNumErr{
			InvalidStatus: InvalidNumOdd,
			message:       fmt.Sprintf("expected an even number received: %d", num),
		}
	}
	return true, nil
}

func main() {
	oddNum := 43
	_, err := isEven(oddNum)

	if err != nil {
		var invalidNumErr InvalidNumErr
		if isInvalidNumErr := errors.As(err, &invalidNumErr); isInvalidNumErr {

			fmt.Printf("Status: %d | Error: %s\n", invalidNumErr.InvalidStatus, invalidNumErr.message)
		} else {
			fmt.Println(err)
		}
	}

}

// errors.As | Compare Error Types

//   - match a returned error and all the errors that it wraps to a specified error type
//   - if an error matches the target error type then the returned error is assigned to the target error
//   - the target must be a pointer to an errortype or a pointer to an interface

// overriding errors.As

//   - can be overridden in a similar way to errors.Is with the additional complexity of reflection
//   - overriding errors.As should only be done in special cases like checking the type of an error and returning a diffrent one

// When to Use errors.As

//   - When you are looking for a specific error type regardless of values
