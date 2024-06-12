package model

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type TcpServer struct {
	WG         sync.WaitGroup
	Listener   net.Listener
	Shutdown   chan struct{}
	Connection chan net.Conn
}

// --------------------- TCP Server Methods ---------------------//

// acceptConnections:
//
// Accepts client connections and passes them into the connection channel.
func (s *TcpServer) acceptConnections() {
	defer s.WG.Done()

	// infinite loop to listen for incoming client connection requests
	for {
		select {
		// close the infinite loop if s.Shutdown is not nil
		case <-s.Shutdown:
			fmt.Println("acceptConnections: done!")
			return

		default:
			conn, err := s.Listener.Accept()

			if err != nil {
				fmt.Println("error when attempting to accept a client request to connect.")
				continue
			}

			fmt.Printf("new client connection accepted from: %s\n", conn.LocalAddr().String())

			// send the accepted client connection into the channel of connections
			s.Connection <- conn
		}
	}
}

// handleConnections:
//
// Recieves new client connections through the connection channel and passes them to a new go routine to be handled.
func (s *TcpServer) handleConnections() {
	defer s.WG.Done()

	for {
		select {

		case <-s.Shutdown:
			fmt.Println("handleConnections: done!")
			return
		// pass accepted client connections to a new go routine to be handled in the order the connections were accepted
		case conn := <-s.Connection:
			go s.handleConnection(conn)

		}
	}
}

// handleConnection:
//
// processes each individual client request.
func (s *TcpServer) handleConnection(conn net.Conn) {
	defer conn.Close()

	buff := make([]byte, 1024)

	_, err := conn.Read(buff)

	if err != nil {
		fmt.Fprint(conn, "There was an error processing your request.\n")
		return
	}

	fmt.Fprintf(conn, "ECHO: %s", buff)

	time.Sleep(2 * time.Second)

	fmt.Fprint(conn, "Closing the connection... Goodbye my friend!\n")
}

func (s *TcpServer) Start() {
	s.WG.Add(2)

	go s.acceptConnections()

	go s.handleConnections()
}

func (s *TcpServer) Stop() {

	close(s.Shutdown)

	s.Listener.Close()

	done := make(chan struct{})

	go func() {
		s.WG.Wait()
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("server shutdown successfully")
		return

	case <-time.After(2 * time.Second):
		fmt.Println("time out waiting for connection to finish.")
		return
	}
}

//--------------------- Functions ---------------------//

// NewTcpServer:
//
// Create a new TCP Server listening on the specified address.
//
// A new socket is created bound to the port provided.
func NewTcpServer(address string) (*TcpServer, error) {
	listener, err := net.Listen("tcp", address)

	if err != nil {
		return nil, fmt.Errorf("failed to listen on address: %s: %w", address, err)
	}

	fmt.Printf("starting TCP Server listening on: %s\n", address)

	return &TcpServer{
		Listener:   listener,
		Shutdown:   make(chan struct{}),
		Connection: make(chan net.Conn),
	}, nil
}
