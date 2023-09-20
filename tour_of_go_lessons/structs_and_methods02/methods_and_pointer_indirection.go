package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(scaler float64) {
	v.X *= scaler
	v.Y *= scaler

}

func ScaleFunc(v *Vertex, scaler float64) {
	v.X *= scaler
	v.Y *= scaler
}

func main() {

	var vector = Vertex{X: 4, Y: 2}

	vectorPtr := &vector

	vector.Scale(10)

	fmt.Println(vector)

	ScaleFunc(vectorPtr, 2)

	fmt.Println(vector)

	(&vector).Scale(3)

	fmt.Println(vector)

	vectorPtr.Scale(5)

	fmt.Println(vector)

}

// Methods and Pointer Indirection

//   - functions with pointer arguments must take pointers

//   - methods with pointer receivers take either
//     a value or a pointer receiver when invoked

//   - v.Scale(5) == (&v).Scale(5) in Go interpreter

//   - the equivalent happens in reverse

//   - functions with value arguments must take a value

//   - methods with value receivers take either
//     a value or a pointer receiver when invoked
