package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// TODO: review an refactor comments

func main() {
	// path to source file
	sourceFilePath := "io_streams/source_data/xanadu.txt"

	destinationFilePath := "io_streams/byte_streams/output.txt"

	// open input source file for reading returning an os.File object
	in, err := os.Open(sourceFilePath)

	if err != nil {
		log.Fatalf("Error opening source file: %s", sourceFilePath)
	}
	// close the file stream to prevent resource leakage
	defer in.Close()

	// create output destination file if it does not exist, truncate the file if it does exist
	out, err := os.Create(destinationFilePath)

	if err != nil {
		log.Fatalf("Error creating destination file: %s", sourceFilePath)
	}

	// close the file stream to prevent resource leakage
	defer out.Close()

	// allocate byte chunk size, in this case we are reading 1 byte at a time where each byte represents is a utf-8 character
	byteSlice := make([]byte, 10)

	for {

		// read data from input file stream into byte slice
		n, err := in.Read(byteSlice)

		// break out of the loop if the end of the file is reached
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error reading source file: %s", sourceFilePath)
		}

		fmt.Printf("n = %v err = %v b = %v\n", n, err, byteSlice)

		fmt.Printf("b = %q\n", byteSlice[:n])

		time.Sleep(500 * time.Millisecond)

		// write data to the destination file stream
		out.Write(byteSlice[:n])

	}

}

// I/O Streams

//   - streams are unidirectional continuous pipelines of sequential data that can read from a source and write to a destination in byte sized  chuncks

//   - I/O Streams represent an input source (reading data) and an output destination (writing data)

//   - there are many types of streams that handle various types of data

//   - input data sources and output data destinations can be anything that holds, generates or consumes data

//   - e.g. files, network connections, network sockets, in-memory buffers, etc

//   - separating input streams into their own interdependent pipelines allows for ease of data flow

// Input Streams

//   - input streams are used to read data from a source

// Output Streams

//   - output streams are used to write data to a destination

// Streams in GO

//  while some languages like Java implement many types of streams go primarily uses byte-oriented streams that are UTF-8 encoded

// Input Streams in GO

//   - io.Reader Implementations

//       - os.File (reading files)

//       - bytes.Buffer (in-memory buffers)

//       - net.Conn (network connection / sockets)

//       - strings.Reader (reading strings)

// io.Writer Implementations

//       - os.File (reading files)

//       - bytes.Buffer (in-memory buffers)

//       - net.Conn (network connection / sockets)

//       - ioutil.Discard (null writer that discards all data)

// Closing Streams and Resource Leakage

// ? What does it mean for I/O operations to be primitive?
// TODO: See https://pkg.go.dev/io for insight on the above question
