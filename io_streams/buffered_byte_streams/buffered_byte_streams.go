package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// TODO: review an refactor comments

func main() {
	// input source filepath
	inputFilePath := "io_streams/source_data/xanadu.txt"

	// output destination filepath
	outputFilePath := "io_streams/buffered_byte_streams/output.txt"

	// open input source file for reading | read-only mode
	in, err := os.Open(inputFilePath)

	if err != nil {
		log.Fatalln(err.Error())
	}

	// close, preventing memory resource leakage
	defer in.Close()

	// create output file if it does not exist, truncate otherwise | read-write mode
	out, err := os.Create(outputFilePath)

	if err != nil {
		log.Fatalln(err.Error())
	}

	// close, preventing memory resource leakage
	defer out.Close()

	// create a buffered reader | input source byte stream
	bufIn := bufio.NewReader(in)

	// create a buffered writer | output destination byte stream
	bufOut := bufio.NewWriter(out)

	// fixedBuffIn := bufio.NewReaderSize(in, 1024)

	for {
		// read string until a newline is encountered
		line, err := bufIn.ReadString('\n')

		// if we reach the end of the file write the final line
		if err == io.EOF {
			_, err := bufOut.WriteString(line)

			if err != nil {
				log.Fatalln(err.Error())
				break
			}

			fmt.Println(line)

			fmt.Println("File copied successfully!")

			break
		}

		if err != nil {
			log.Fatalln(err.Error())
		}
		// write line to output buffer
		_, err = bufOut.WriteString(line)

		if err != nil {
			log.Fatalln(err.Error())
		}

		fmt.Println(line)
	}
	// write data from output internal buffer to output destination
	bufOut.Flush()

}

// Buffered Byte Streams

//   - a wrapper around the io.Reader and io.Writer interfaces that add an intermediate internal buffer

//   - bufio.Reader and bufio.Writer also have additional helper methods such as ReadSring and WriteString

// How Buffered Byte Streams Work Under The Hood

//   - Internal Buffer:

//       ~ both the bufio.Reader and bufio.Writer are created with an internal default buffer
//         typically an array of bytes 4kb is size

//   - Reading Data From a Buffered Byte Stream

//       ~ the buffer is checked for data, if the buffer is empty then data will be read from the underlying byte stream into the buffer

//   - Writing Data to a Buffered Byte Stream

//       ~ when data is written to a buffered byte stream the data is stored in the buffer and will not be written to the output destination

//       ~ the output destination byte stream is written to only when the buffer is filled or the buffer is explicitly flushed

// Why Use Buffered Byte Streams

//   - Performance Improvement

//       ~ reduction in the amount of system calls made to the underlying byte stream (e.g. file or network socket) by batching operations

//       ~
