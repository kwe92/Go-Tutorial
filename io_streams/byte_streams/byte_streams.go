package main

import (
	"log"
	"os"
)

func main() {
	sourceFilePath := "io_streams/byte_streams/xanadu.txt"

	inputSourceFile, err := os.Open(sourceFilePath)

	if err != nil {
		log.Fatalf("Error opening source file: %s", sourceFilePath)
	}

	defer inputSourceFile.Close()

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
