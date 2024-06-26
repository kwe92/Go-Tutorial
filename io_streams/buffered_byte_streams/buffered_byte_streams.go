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
		switch {
		case err == io.EOF:
			_, err := bufOut.WriteString(line)

			if err != nil {
				log.Fatalln(err.Error())
				continue
			}

			fmt.Println(line)

			// write data from output internal buffer to output destination
			bufOut.Flush()

			fmt.Println("File copied successfully!")

			return

		case err != nil:
			log.Fatalln(err.Error())
		}

		// write line to output buffer
		_, err = bufOut.WriteString(line)

		if err != nil {
			log.Fatalln(err.Error())
		}

		fmt.Println(line)
	}

}

// Buffered Byte Streams

//   - a wrapper around the io.Reader and io.Writer interfaces that add an intermediate internal temporary buffer in memory
//     and optimizes performance of file I/O and network operations

//   - bufio.Reader and bufio.Writer also have additional helper methods such as ReadSring and WriteString

// How Buffered Byte Streams Work Under The Hood

//   - Convert Unbuffered Streams to Buffered Streams

//       ~ unbuffered byte streams are passed to the buffered streams to be wrapped

//   - Internal Buffer:

//       ~ both the bufio.Reader and bufio.Writer are created with an internal default buffer
//         typically an array of bytes 4kb is size

//   - Reading Data From a Buffered Byte Stream

//       ~ the buffer is checked for data, if the buffer is empty then data will be read from the underlying byte stream into the buffer

//   - Writing Data to a Buffered Byte Stream

//       ~ when data is written to a buffered byte stream the data is stored in the buffer and will not be written to the output destination

//   - Flushing Buffered Output Streams

//       ~ flushing is the process of writing data from an output buffer to its destination

//       ~ flushing can occur under two conditions in GO, when the buffer is filled or the buffer is explicitly flushed with its flush method

// Why Use Buffered Byte Streams

//   - Performance Improvement

//       ~ reduction in the amount of system calls (random access operations) made to the underlying byte stream (e.g. file or network socket) by batching operations
//         leading to a reduction in network calls and operating system reading and writing data to disk when is typically expensive

//       ~ the reduction in system csalls also increases data locality

//       ~ buffered readers can pre-fetch data making read operations faster

//       ~ buffered writers can buffer data before writing to the file, reducing the number of disk writes
//         this is why flushing is required when working with buffered byte streams

//       ~ custimizable buffer size when creating bufio.Reader and bufio.Writer
