package main

import "fmt"

// Interface Type

//   - a set of method signatures
//   - similar to an abstract class in other object oriented languages
//   - no concrete implementation is needed for Shape interface itsself.

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

// implementation of Area method for the Circle struct

func (c *Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func PrintArea(s Shape) {
	fmt.Printf("\nArea: %.2f\n\n", s.Area())
}

func main() {
	c := Circle{Radius: 6}

	s := Shape(&c)

	PrintArea(s)
}
