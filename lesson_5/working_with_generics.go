package main

import "fmt"

// similar to an abstract interface, no concrete implementation is needed for Shape struct itsself.

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

// implementation of Area method for the Circle struct

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func PrintArea(s Shape) {
	fmt.Println("\nArea:", s.Area(), "\n")
}

func main() {
	c := Circle{Radius: 6}

	PrintArea(c)
}
