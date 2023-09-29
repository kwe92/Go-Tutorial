package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	url := "https://vast-puce-mite-fez.cyclic.app/animeme"

	response0, err := http.Get(url)

	checkError(err)

	//? Unmarshal / Decode io.ReadCloser With Custom Implementation

	fmt.Printf("\nresponse:\n\n%+v\n\n", response0)

	fmt.Printf("\nresponse body:\n\n%+v\n\n", response0.Body)

	buf := make([]byte, response0.ContentLength)

	ReadAndClose(response0.Body, buf)

	fmt.Printf("\nresponse body as Slice of bytes:\n\n%+v\n\n", buf)

	fmt.Printf("\nresponse body as string:\n\n%+v\n\n", string(buf))

	//? Unmarshal / Decode io.ReadCloser With ioutils Helper Package

	// response1, err := http.Get(url)

}

func ReadAndClose(r io.ReadCloser, buf []byte) (n int, err error) {

	for len(buf) > 0 {

		var nr int

		if err == io.EOF {
			return
		}

		nr, err = r.Read(buf)

		n += nr

		buf = buf[:nr]

		fmt.Println("\nbuffer bytes from ReadAndClose: ", buf[:n])

		fmt.Println(string(buf))

	}
	r.Close()
	return
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// http.Response

//   - a struct defined in the http package that represents an HTTP response0 received
//     for a request

// response0.Body | io.ReadCloser

//   - the content of an HTTP response0 is an io.ReadCloser implementation
//   - the content can not be represented as a string until it is unmarshalled or decoded
//   - there are several ways to unmarshal an io.ReadCloser

// Least Efficient Ways:

// Implement ReadAndClose function

//   - which takes an io.ReadCloser and a Slice of bytes as an argument
//   - and continuously reads into the Slice of bytes until the end of the file reached
//   - the Slice of bytes can then be converted to a string with type conversion

// Use ioutils helper package

//   - pass the io.ReadCloser as an argument to ioutils.ReadAll
//     which returns a Slice of bytes the length of the content and any errors encountered
//   - the Slice of bytes can then be converted to a string with type conversion
