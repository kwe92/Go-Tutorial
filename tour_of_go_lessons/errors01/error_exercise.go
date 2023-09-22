package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(Sqrt(2))

	fmt.Println(Sqrt(-2))

}

// Sqrt a root-finding algorithm using newton's method
func Sqrt(x float64) (float64, error) {

	if x < 0 {

		return 0, ErrNegativeSqrt(x)

	}
	z := 1.0

	for i := 0; i < 10; i++ {

		z -= (z*z - x) / (2 * z)

		fmt.Println(z)

	}

	return z, nil
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {

	errorTimeStamp := time.Now()

	return fmt.Sprintf(

		"\n At: %v | Error: can't Sqrt negative number: %v",

		errorTimeStamp,

		float64(e),
	)
}

// Summary

//   - implement a square root function that
//     returns the square root of a non-negative rational number
//   - if the number passed in is negative then return the zero value
//     and the error type ErrNegativeSqrt using a type conversion
//   - define a type ErrNegativeSqrt that is a float in order to pass
//     in the negative number as an argument to be printed
//   - implement an Error method with a value receiver that
//     prints the string representation of the error you would like to present
