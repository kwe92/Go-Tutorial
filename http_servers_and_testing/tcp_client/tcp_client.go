package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

// Creating a Client Using Sockets

var (
	PROTOCOL = "tcp"            // communication protocol: Transmission Control Protocol
	HOSTNAME = "127.0.0.1:8080" // server and port number you want to connect to
)

func main() {

	for i := 0; i < 5; i++ {
		time.Sleep(2000 * time.Millisecond)

		// create a new socket of type TCP bound to port 8080 requesting communcation to a remote server
		// if a connection is successfully established then the client can read from and write to the socket
		serverConnection, err := net.Dial(PROTOCOL, HOSTNAME)

		if err != nil {
			log.Fatalf("error when trying to establish a connection to %s\n", HOSTNAME)
		}

		// attempt to write data output to the remote server over the socket
		_, err = serverConnection.Write([]byte("This is a message from the Go Client!"))

		if err != nil {
			log.Fatalf("error when trying to write data to the server at: %s\n", HOSTNAME)
		}

		// create buffer to read incoming bytes of data
		buff := make([]byte, 1024)

		// attempt to read incoming remote server bytes of data as input into the socket
		serverConnection.Read(buff)

		if err != nil {
			fmt.Printf("error reading data from server: %s\n", err.Error())
			return
		}

		fmt.Printf("data recieved from server connection: %s\n", string(buff))

		serverConnection.Close()

	}

}

// Establishing a client connection to a server using sockets

//   - Client needs to know:

// 		 - the protocol used for communication (i.e. TCP or UDP)

//       - the host name (IP address) of the machine

//       - the port number the server is listening on\

// Port Numbers Range

//   - ports are in the range of uint16, a sub-set of whole numbers ranging from 0 to 65,535 i.e. 2 ** 16

// TCP Client Side:

//   - writes to the TCP Server and then waits for a response
