package main

import (
	"fmt"
	"math"
)

// type signature for a function
type FloatCallback func(float64, float64) float64

type Vertex struct {
	X, Y float64
}

// functions are objects and can be passed as arguments then invoked in separate parts of your program

func compute(fn FloatCallback, v *Vertex) float64 {
	return fn(v.X, v.Y)
}

var hypot FloatCallback = func(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func main() {

	v := Vertex{X: 3, Y: 4}

	fmt.Println(hypot(6, 8))

	fmt.Println(compute(hypot, &v))

	fmt.Println(compute(math.Pow, &v))

}
