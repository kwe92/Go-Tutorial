//  TODO: Edit Comments

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	var err error

	var data []byte

	var alphabetData []byte

	var someFile string = "/Users/kwe/GoLang/go-Tutorial/tour_of_go_lessons/reader01/some_text_file.txt"

	var alphbetFile string = "/Users/kwe/GoLang/go-Tutorial/tour_of_go_lessons/reader01/alphabet.txt"

	// read file, return Slice of bytes
	data, err = os.ReadFile(someFile)

	checkError(err)

	// read file, return Slice of bytes
	alphabetData, err = os.ReadFile(alphbetFile)

	checkError(err)

	// convert byte Slice to string

	var string_rep string = bytesToString(data)

	// convert bytes Slice to string

	var alpha_string string = bytesToString(alphabetData)

	fmt.Println("\ndata as byte Slice:", data)

	fmt.Println("\ndata string representation:", string(string_rep))

	fmt.Println("\nalphabet as byte Slice:", alphabetData)

	fmt.Println("\nalphabet string representation:", string(alpha_string))

	fmt.Printf("bytes: %d | chars: %d\n\n", len(alphabetData), len(alpha_string))

	// Open file for reading
	file, err := os.Open(someFile)

	checkError(err)

	// allocate and initialize Slice of bytes with a size of 8

	b0 := make([]byte, 8)

	// read file 8 bytes at a time (the length of b0, our byte Slice)
	for {
		n0, err := file.Read(b0)

		// if the err indicates you are at the end of the file break the loop
		if err == io.EOF {
			fmt.Printf("\n\n%v\n\n", err)
			break
		}

		fmt.Printf("\nn = %v\n\nb = %v\n\nstr = %q\n\nerr = %v", n0, b0, string(b0[:n0]), err)

	}

	// Seek to a known location within a file, skipping elements

	offset := int64(6)

	// Set the offset for the next read of a file
	o0, err := file.Seek(offset, 0)

	checkError(err)

	b1 := make([]byte, 4)

	n0, err := file.Read(b1)

	checkError(err)

	fmt.Printf("%d bytes @ %d: ", n0, o0)

	fmt.Printf("%v\n", string(b1[:n0]))

	o1, err := file.Seek(offset, 0)

	checkError(err)

	b2 := make([]byte, 4)

	n1, err := io.ReadAtLeast(file, b2, 1)

	checkError(err)

	fmt.Printf("%d bytes @ %d: %s\n", n1, o1, bytesToString(b2))

	// Rewinding Using Seek(0, 0)

	// return to the beginning of a file
	_, err = file.Seek(0, 0)

	checkError(err)

	reader := bufio.NewReader(file)

	b3, err := reader.Peek(6)

	fmt.Printf("6 bytes: %s\n", bytesToString(b3))

	// close the file when you are done
	file.Close()

}

// checkError checks for an error, stops execution if an error evaluates to not null and prints the error.
func checkError(err error) {

	if err != nil {
		panic(err)
	}

}

// bytesToString converts a Slice of bytes into a string representation.
func bytesToString(bytes []byte) string {
	return string(bytes)
}

// os.ReadFile

//   - ReadFile is part of the `os` package provided the Go standard library
//   - ReadFile reads a file and return the files contents as a Slice of bytes

// bufio

//   - implements a buffered reader
//   - efficient for many small reads
//   - provides additional reading methods
