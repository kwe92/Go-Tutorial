package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

// Creating a Client Using Sockets

var (
	PROTOCOL = "tcp"            // the agreed upon communication protocol : Transmission Control Protocol
	HOSTNAME = "127.0.0.1:8080" // the server and port number you want to connect to
)

func main() {

	for i := 0; i < 5; i++ {
		time.Sleep(2000 * time.Millisecond)

		// request a connection to the server
		serverConnection, err := net.Dial(PROTOCOL, HOSTNAME)

		if err != nil {
			log.Fatalf("error when trying to establish a connection to %s\n", HOSTNAME)
		}

		// attempt to write data to the server
		_, err = serverConnection.Write([]byte("This is a message from the Go Client!"))

		if err != nil {
			log.Fatalf("error when trying to write data to the server at: %s\n", HOSTNAME)
		}

		// create a buffer to read incoming data
		buff := make([]byte, 1024)

		serverConnection.Read(buff)

		if err != nil {
			fmt.Printf("error reeading data from client connection in handleClientConnection: %s\n", err.Error())
			return
		}

		fmt.Printf("data recieved from server connection: %s\n", buff)

		serverConnection.Close()

	}

}

// Establishing a client connection to a server using sockets

//   - Client needs to know:

//       - the host name (IP address) of the machine

//       - the port number the server is listening on

// Port Numbers Range

//   - ports are in the range of uint16, a sub-set of whole numbers ranging from 0 to 65,535
