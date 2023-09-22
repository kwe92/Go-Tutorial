//  TODO: add comments

package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("In the beginning was the word.")

	b := make([]byte, 64)

	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}

	}
}
