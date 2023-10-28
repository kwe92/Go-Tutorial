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
//   - the type of receiver implemented dictates if a type received is mutable or immutable

// Value Receivers

//   - value receivers are copies and the original object remains immutable by the method
//   - copies of objects allocate new memory for that object making value receivers less efficient
//   - can not work with nil receivers
//   - only contains the method set for value receivers
//     if a method is implemented with a pointer receiver then a pointer is required
//     as the receiver will not show wup in the value receiver method set

// Pointer Receivers

//   - pointer receivers point to an object in memory rendering them mutable to a method
//   - preferred over value receivers as they are more memory efficient
//   - contains the method set for both value and pointer receivers
//     implying that a pointer can be used anywhere a value is expected
