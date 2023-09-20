package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {

	vectorAbs := math.Sqrt(v.X*v.X + v.Y*v.Y)

	return vectorAbs
}

func (v *Vertex) Scale(scaler float64) {
	v.X *= scaler
	v.Y *= scaler
}

func main() {
	vector := Vertex{3, 4}

	fmt.Println(vector.Abs())

	vector.Scale(10)

	fmt.Println(vector)

	fmt.Println(vector.Abs())

}

// Method Receiver Types

//   - methods can be declared with value receivers or pointer receivers
//   - pointer receivers are typically chosen over value receivers

// Value Receivers

//   - value receivers are copies and the original object is rendered immutable by a method
//   - copies of objects also allocate new memory for that object making value receivers less efficient

// Pointer Receivers

//   - pointer receivers point to an object in memory rendering them muttable to a method
//   - preffered over value receivers as they are more memory efficient
