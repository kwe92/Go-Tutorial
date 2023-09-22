//  TODO: add comments

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	var data []byte

	var err error

	var file string = "/Users/kwe/GoLang/go-Tutorial/tour_of_go_lessons/reader01/some_text_file.txt"

	var alphbetFile string = "/Users/kwe/GoLang/go-Tutorial/tour_of_go_lessons/reader01/alphabet.txt"

	data, err = os.ReadFile(file)

	check(err)

	alphaByte, err := os.ReadFile(alphbetFile)

	check(err)

	var string_rep string = bytesToString(data)

	var alpha_string string = bytesToString(alphaByte)

	fmt.Println(data)

	fmt.Println(string(string_rep))

	fmt.Println(alphaByte)

	fmt.Println(string(alpha_string))

	fmt.Printf("bytes: %d | chars: %v\n\n", len(alphaByte), len(alpha_string))

	f, err := os.Open(file)

	check(err)

	b0 := make([]byte, 8)

	for {
		n0, err := f.Read(b0)

		if err == io.EOF {
			fmt.Printf("\n\n%v\n\n", err)
			break
		}

		fmt.Printf("\nn = %v\n\nb = %v\n\nstr = %q\n\nerr = %v", n0, b0, string(b0[:n0]), err)

	}

}

func check(err error) {

	if err != nil {
		panic(err)
	}

}

func bytesToString(bytes []byte) string {
	return string(bytes)
}
