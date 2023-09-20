package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

var vector Vertex

var float = MyFloat(-3.14)

var i Abser

func main() {

	vector = Vertex{6, 8}

	vectorPtr := &vector

	i = float

	fmt.Println(i.Abs())

	i = vectorPtr

	fmt.Println(i.Abs())

	// throws an error at compile-time
	// i = vector
}

type Vertex struct {
	X, Y float64
}

type MyFloat float64

func (v *Vertex) Abs() float64 {
	vectorAbs := math.Sqrt(v.X*v.X + v.Y*v.Y)

	return vectorAbs
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// GO Interface Type

//   - interfaces are a type in GO defined as a set of method signatures
//   - any value `a defined type` that implements the methods of an interface
//     can be a value of that interface type

// Structual Typed Interfaces

//   - analogous to duck typed interfaces but errors are caught at compile-time instead of run-time `tooling in the IDE of choice`
//   - Interface values must at least implement the method signatures of an interface type but can have additional methods

// Type Checking Interfaces

//   - can be statically checked via the GO type system at compile-time
//   - can be dynamically checked with type assertions

// Interfaces and Method Receiver Types

//   - if a method uses a pointer receiver implementation a pointer must be used as a value to the implemented interface
//   - if a method uses a value receiver implementation a pointer or a value receiver can be used as a value to the implemented interface
