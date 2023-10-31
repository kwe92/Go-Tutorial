package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	vectorAbs := math.Sqrt(v.X*v.X + v.Y*v.Y)

	return vectorAbs
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {

	var vector Vertex

	var float = MyFloat(-3.14)

	var abser Abser

	vector = Vertex{6, 8}

	vectorPtr := &vector

	abser = float

	fmt.Println(abser.Abs())

	abser = vectorPtr

	fmt.Println(abser.Abs())

	// throws an error at compile-time
	// i = vector
}

// Program to Interfaces and Not Implementations - The Gang of Four

//   - Depend on behavior and not concrete implementations

// GO Interface Type

//   - interfaces are a type in GO defined as a set of method signatures
//   - they are the only abstract type in GO similar to abstract classes in other languages
//   - any built-in type or user defined type that implements the methods of an interface
//     can be a value of that interface type
//   - interface names typically end with `er` e.g. Abser

// Structual Typed Interfaces

//   - type safe duck typing
//   - analogous to duck typed interfaces but errors are caught at compile-time instead of run-time `tooling in the IDE of choice`
//   - Interface values must at least implement the method signatures of an interface type but can have additional methods

// Type Checking Interfaces

//   - can be statically checked via the GO type system at compile-time
//   - can be dynamically checked with type assertions

// Interfaces and Method Receiver Types

//   - if a method uses a pointer receiver implementation a pointer must be used as a value to the implemented interface
//   - if a method uses a value receiver implementation a pointer or a value receiver can be used as a value to the implemented interface
