// Exercise: Readers

// Implement a Reader type that emits an infinite stream of the ASCII character 'A'.

package main

import (
	"fmt"
)

func main() {

	// declare and initialize a Slice of bytes that has a size of 8
	byteSlice := make([]byte, 8)

	fmt.Println("\nSlice of 8 bytes initialized with their zero value: ", byteSlice)

	// declare and initialize the MyReader struct
	reader := MyReader{}

	// invoke the Read method passing in the slice of bytes as an argument
	n, err := reader.Read(byteSlice)

	fmt.Printf("\nn = %v, err = %v\n\n", n, err)

}

type MyReader struct{}

// Read implements the io.Reader interface.
func (r MyReader) Read(b []byte) (n int, err error) {

	for i := range b {
		b[i] = 65
	}

	fmt.Printf("\n%v\n", b)

	n = len(b)

	return
}

// Summary

//   - implement a method called Read with the signature of Read([]byte) (int, error)
//     on the MyReader structure type defined

//   - implementing Read implicitly implements the io.Reader interface

//   - the implementation of Read on the defined type MyReader should emit a stream of the
//     ASCII character 'A'

//   - which implies assgining the byte representation of `A` to each element of the Slice of bytes
//     passed in as an argument until all elements represent the byte representation of the ASCII character `A`
