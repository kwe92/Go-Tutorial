package main

import (
	"fmt"
	"log"
	"net"
)

// Creating a Echo Server Using Sockets

var (
	PROTOCOL = "tcp" // use the Transmission Control Protocol
	PORT     = ":8080"
)

func main() {
	// create a new socket bound to port 8080
	// listening for an incoming connection from a client
	serverListener, err := net.Listen(PROTOCOL, PORT)

	if err != nil {
		log.Fatalf("There was an error when establishing the server TCP connection: %s\n", err.Error())
	}

	fmt.Printf("TCP server started successfully, listening on port %s\n", PORT)

	// run server in infinite while loop
	for {

		// Accept incoming client connections
		ClientConn, err := serverListener.Accept()

		if err != nil {
			log.Printf("error with new client connection: %s\n", err.Error())
			continue
		}

		// launch a new Go routine to handle multiple client connections to the server
		go handleClientConnection(ClientConn)

	}

}

func handleClientConnection(ClientConn net.Conn) {
	// close the connection to the client when finished

	// defer ClientConn.Close()

	// create a buffer to read incoming data
	buff := make([]byte, 1024)

	_, err := ClientConn.Read(buff)

	if err != nil {
		fmt.Printf("error reeading data from client connection in handleClientConnection: %s\n", err.Error())
		return
	}

	fmt.Printf("data recieved from client connection: %s\n", buff)

	ClientConn.Write(buff)
}

// What is a Socket?

//   - one endpoint of a two-way communication link between two programs running on a network

//   - a socket is bound to a port number allowing the TCP layer to identify
//     the application that the data is destined to be sent to for both client requests and server responses

// Identifying TCP Connections

//   - a TCP connection can be uniquely identified by
//   - the composition of its two endpoints i.e. IP address and port number

// Socket programming is an important skill for network programming

// building powerful network applications with socket programming and the Client-Server Model
