// TODO: add explanitory comments

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {

	// anime meme API URL
	url := "https://vast-puce-mite-fez.cyclic.app/animeme"

	// http.Client instance with timeout set
	client := http.Client{Timeout: 10 * time.Second}

	// perform HTTP GET request with instantiated Client
	response0, err := client.Get(url)

	// check errors received from request
	checkError(err)

	//? Read content from io.ReadCloser With Custom Implementation

	// print response0 struct to the console
	fmt.Printf("\nresponse0:\n\n%+v\n\n", response0)

	// print response0 Body property to the console
	fmt.Printf("\nresponse0 body:\n\n%+v\n\n", response0.Body)

	// allocate and initalize a Slice of bytes the length of the response0 Body contents
	response0Bytes := make([]byte, response0.ContentLength)

	// read the contents of the response0 Body into the response0Bytes byte Slice allocated
	ReadAndClose(response0.Body, response0Bytes)

	// print the contents as a Slice of bytes to the console
	fmt.Printf("\nresponse0 body as Slice of bytes:\n\n%+v\n\n", response0Bytes)

	// print the contents as a string to the console
	fmt.Printf("\nresponse0 body as string:\n\n%+v\n\n", string(response0Bytes))

	//? Read content from io.ReadCloser With io Package

	response1, err := client.Get(url)

	// instead of allocating a separate Slice and then reading into that Slice we do it simultaniously with io.ReadAll
	response1Bytes, err := io.ReadAll(response1.Body)

	checkError(err)

	// convert Slice of bytes to a string representation
	response1StringRep := string(response1Bytes)

	// print the converted string representation to the console
	fmt.Printf("\nresponse1 string representation:\n\n%+v\n\n", response1StringRep)

	//? Read content from io.ReadCloser with bytes package

	response2, err := client.Get(url)

	// allocate and initalize a Slice of bytes the length of the response2 contents
	response2ByteSlice := make([]byte, response2.ContentLength)

	// create and initialize a lebuffer passing in the allocated Slice the length of the contents
	buffer := bytes.NewBuffer(response2ByteSlice)

	// read the contents of the io.ReadCloser into the buffer
	_, err = buffer.ReadFrom(response2.Body)

	checkError(err)

	// return a string representation of the buffer
	response2StringRep := buffer.String()

	// print the string representation tothe console
	fmt.Printf("\nresponse2 string representation:\n\n%+v\n\n", response2StringRep)

	//? Read content from io.ReadCloser with json package

	// define a struct the shape of the request
	type AnimeMeme struct {
		Title string `json:"title"`
		Url   string `json:"url"`
	}

	var jsonData any

	var animeMeme AnimeMeme

	response3, err := client.Get(url)

	checkError(err)

	response4, err := client.Get(url)

	checkError(err)

	// use json.NewDecoder to buffer io.ReadCloser and Decode to store its contents in a type, seems to default to a map
	json.NewDecoder(response3.Body).Decode(&jsonData)

	// Decode the response into a struct
	json.NewDecoder(response4.Body).Decode(&animeMeme)

	// print the default data structure the data is loaded into
	fmt.Printf("\nresponse3 GO data structure representation:\n\n%+v\n\n", jsonData)

	// print the struct type defined above with data written from the Decode process
	fmt.Printf("\nresponse3 GO struct representation:\n\n%+v\n\n", animeMeme)

}

// ReadAndClose continuously reads the contents of r into buf until the end of the file is reached
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

// Ways to Read content from an io.ReadCloser | Worst to Best:

// Implement ReadAndClose function

//   - which takes an io.ReadCloser and a Slice of bytes as an argument
//   - and continuously reads into the Slice of bytes until the end of the file reached
//   - the Slice of bytes can then be converted to a string with type conversion

// Use io package

//   - pass the io.ReadCloser as an argument to io.ReadAll
//     which returns a Slice of bytes the length of the content and any errors encountered
//   - the Slice of bytes can then be converted to a string with type conversion

// Use bytes package

//   - write content of an io.ReadCloser into an in-memory buffer that grows to the size of the content

// Use json Package

//   - Decode / unmarshal the response into a target variable
