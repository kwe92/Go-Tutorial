package main

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

type ReadCloser interface {
	Reader
	Closer
}

func main() {}

// Embedded Interface

//   - As you can embed a type in a struct you can embed an interface within an interface
