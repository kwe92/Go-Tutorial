package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	inputSourceFilePath := "io_streams/source_data/xanadu.txt"

	outputDestinationFilePath := "io_streams/scanners/output.txt"

	in, err := os.Open(inputSourceFilePath)

	checkErr(err)

	defer in.Close()

	out, err := os.Create(outputDestinationFilePath)

	checkErr(err)

	defer out.Close()

	scanner := bufio.NewScanner(in)

	// scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println(line)

		_, err := out.WriteString(line + "\n")

		checkErr(err)

	}

	checkErr(scanner.Err())

}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

// Scanner

//   - a tool used to efficiently read and parse text from input streams

//   - when passed an unbuffered input stream a Scanner created a buffer byte stream under the hood similar to bufio.NewReader

//   - has similar benifits to buffered byte streams produced by bufio.NewReader with less complexity involved

//   - data is tokenized and can be read one line at a time, by word, or by whatever delimiter you want

//   - automatically change the buffer size when input is large

//   - best used for text files that are read by some delimiter
