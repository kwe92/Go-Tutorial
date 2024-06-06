package main

import (
	"fmt"
	"log"
	"net"
)

// Creating a Echo Server Using Sockets

var (
	PROTOCOL = "tcp" // communication protocol: Transmission Control Protocol
	PORT     = ":8080"
)

func main() {
	// create a new socket of type TCP bound to port 8080
	// listening for incoming connection requests from clients
	// the socket stays constant and does not change
	serverListener, err := net.Listen(PROTOCOL, PORT)

	if err != nil {
		log.Fatalf("There was an error when establishing the server TCP connection: %s\n", err.Error())
	}

	fmt.Printf("TCP server started successfully, listening on port %s\n", PORT)

	// run server in infinite while loop / start server
	for {

		// accept incoming client connection
		// the server creates a new socket to communicate with
		// the new socket has its remote IP address and remote port set to that of the clients
		ClientConn, err := serverListener.Accept()

		if err != nil {
			log.Printf("error with new client connection: %s\n", err.Error())
			continue
		}

		// launch a new Go routine (thread) to handle multiple client connections (sockets) to the server
		go handleClientConnection(ClientConn)

	}

}

func handleClientConnection(ClientConn net.Conn) {
	// close the connection to the client when finished

	// create a buffer to read incoming data
	buff := make([]byte, 1024)

	// socket input stream / read incoming request data
	_, err := ClientConn.Read(buff)

	if err != nil {
		fmt.Printf("error reading data from client connection in handleClientConnection: %s\n", err.Error())
		return
	}

	fmt.Printf("data recieved from client connection: %s\n", buff)

	// socket output stream
	ClientConn.Write([]byte(fmt.Sprintf("ECHO: %s", string(buff))))
}

// Transmissino Control Protocol

//   - a connection based protocol used to server consistant reliable ordered communication between two programs

//   - built in error handling to ensure all data that was sent was recieved

// TCP Server

//   - Listens for incoming TCP connection requests (from clients)

//   - upon recieving a request a connection to the client is established through a communication channel (stream)

//   - depending on the type of request the server responds accordingly (this is where route handling becomes useful)

//   - the server typically maintains and closes the connections (socket) when communication is finished

// What is a Socket?

//   - one endpoint of a two-way communication link between two programs running on a network

//   - a socket is bound to a port number

//   - a socket is uniquely identified by the composition of IP address and port number

//   - a sockets unique identifier ensures that communication is happening between the server and the correct client

// Components of a Socket

//   - Each socket consists of:

//       - IP address: Identifies the device on the network.

//       - Port number: Specifies the specific application or service on that device.

//       - Protocol: Defines the rules of communication (e.g., TCP (ordered, reliable, consistent, and connection oriented) or UDP (unreliable, potentially unordered, sent in packets that can droped, connectionless)).

// How Sockets Work

//   - Socket Creation:

//       - upon creation a socket is given a protocol type typically TCP or UDP

//   - Binding

//       - a socket is bound to a specified port and IP address on the machine in which it was created

//   - Connecting TCP Sockets

//       - Connections are estiablished with the remote socket (server socket)

//   - Listening on TCP Sockets

//       - servers listen on TCP sockets for incoming connection requests

//   - Data Transfer

//       - data can be transmitted as a stream of bytes between both applications

//   - Closing Sockets

//       - when the communication between both parties is finished sockets should be closed as to not leak resources

// Ports vs Sockets:

//   - Purpose:

//       - Ports: identify applications

//       - Sockets: enable communication between applications

//   - Type:

//       - Ports: are numerical identifiers

//       - Sockets: are complex structures with multiple elements

//   - Scope:

//       - Ports: are local to a computer

//       - Sockets: involve two endpoints on potentially different devices

// Identifying TCP Connections

//   - a TCP connection can be uniquely identified by
//     the composition of its two endpoints i.e. IP address and port number

// Building powerful network applications with socket programming and the Client-Server Model

//   - socket programming is an important skill for network programming

// Server Side Socket Binding

//   - servers run on a specific machine identified by its hostname / ip address

//   - sockets are then bound to a port number specified by some listener or mux implementation

//   - servers listen to the socket for incoming connection requests from clients

//   - the server then determines how to deal with the request

//   - if a server accepts a clients request to connect then a new socket is bound to the same
//     local port as the client requesting services

//   - the remote endpoint (socket) is set to the IP address and port of the client

//   - the afromentioned process allows servers to concurrently process requests by always listen for incoming client
//     requests to connect on the original listening socket

// Client Server Applications

//   - servers provided a service such as processing database queries or transmitting real-time information

//   - clients use the services provided by a server to display information to the user
//     be it some successful indication of an action or changing some mutable state on the client side

// ? What Are The Component Parts That Allow A Client and a Server to Communicate?
