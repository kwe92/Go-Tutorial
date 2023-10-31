package main

import (
	"fmt"
)

// Aiming to Accomplish

//   - Create a way to configure your libraries with default and optional parameters

type Server struct {
	id      string
	maxConn int
	tls     bool
}

// Opts contain options for configuring your server.
type Opts struct {
	id      string
	maxConn int
	tls     bool
}

// ServerWithOpts has a single property opts to specify configuration options.
type ServerWithOpts struct {
	opts Opts
}

// OptsFunc is a function signature for an callback that takes a Opts pointer as an argument.
type OptsFunc func(*Opts)

// defaultOpts returns default Opts that can be modified.
func defaultOpts() Opts {
	return Opts{
		id:      "default",
		maxConn: 10,
		tls:     false,
	}
}

// withTLS is a callback that takes a pointer to Opts struct.
func withTLS(opts *Opts) {
	opts.tls = true
}

// withMaxConn returns a callback that takes a pointer to Opts struct.
func withMaxConn(n int) OptsFunc {
	return func(opts *Opts) {
		opts.maxConn = n
	}
}

// newServer takes n arguments needed to configure a server.
func newServer(
	id string,
	maxConn int,
	tls bool,
) *Server {
	return &Server{
		id:      id,
		maxConn: maxConn,
		tls:     tls,
	}
}

// newServerWithOpts takes a singe argument Opts to configure a server.
func newServerWithOpts(opts Opts) *ServerWithOpts {
	return &ServerWithOpts{
		opts: opts,
	}
}

// TODO: review ...OptsFunc | review the whole function

// newServerWithDefaultOpts provides default configuration options and can take N number of OptsFunc callback to modify default Opts.
func newServerWithDefaultOpts(opts ...OptsFunc) *ServerWithOpts {

	defaultOpts := defaultOpts()

	// enumerate and call any OptsFunc's passed in as arguments
	for _, fn := range opts {
		fn(&defaultOpts)
	}
	return &ServerWithOpts{
		opts: defaultOpts,
	}
}

func main() {
	opts := Opts{
		id:      "V9999",
		maxConn: 2,
		tls:     true,
	}

	// Default way of calling a server
	server := newServer("fx1001", 3, false)

	// Second best way of calling a server with struct
	serverWithOpts := newServerWithOpts(opts)

	// Optimal way if you want to provide optional default arguments
	serverWithDefault := newServerWithDefaultOpts(withTLS, withMaxConn(999))

	fmt.Println("\n", server)

	fmt.Println("\n", serverWithOpts)

	fmt.Printf("\n%+v", serverWithDefault)

}
