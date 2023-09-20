package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

func (v Vertex) abs() float64 {
	vectorAbs := math.Sqrt(v.X*v.X + v.Y*v.Y)
	return vectorAbs
}

func (f MyFloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// receiver prefix ommited, struct passed in as an argument
func abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

var vertex = Vertex{X: 3, Y: 4}

var float MyFloat

func main() {

	float = 3.142

	fmt.Println("\nvertex absolute value from method: ", vertex.abs())

	fmt.Printf("\nfloat absolute value: %.3f", float.abs())

	fmt.Printf("\n\nvertex absolute value from function: %v\n\n", abs(vertex))
}

// Methods && Types

//   - Methods: functions with receiver arguments
//   - Methods can be attached to any type that can be defined:
//       ~ (structs, slices, arrays, floats, maps, etc)
//   - similar to extension methods in Dart with a caveat:
//       ~ the type definition and attached method implementation must be in the same package

// Method Receiver

//   - Prefixes function name
//   - contains variable and type of attached method
//   - example (The Cerberus Function):
//       ~ func (t DefinedType) method_name(param0 data_type) (return_data_type,  return_data_type) {...}
//
