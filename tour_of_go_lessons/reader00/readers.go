//  TODO: add comments

package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("In the beginning was the word.")

	b := make([]byte, 8)

	for {
		n, err := r.Read(b)

		if err == io.EOF {
			break
		}

		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)

		fmt.Printf("b[:n] = %q\n", b[:n])

	}
}

// io.Reader | type Reader interface { Read(p []byte) (n int, err error) }

//   - defined in the io package part of the Go standard library and wraps the Read method signature

// Read Method | Read(p []byte) (n int, err error)

//   - reads up to the lenth of `p` bytes into `p`
//   - where `p` is a Slice of bytes of some length greater than 0 initialized with zero-values
//   - the variable `p` is usually implemented as `b`
//   - returns the number of bytes read and any error encountered
//   - once the stream of bytes is completly read the io.EOF `End of File` error is returned
