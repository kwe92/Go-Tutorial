package main

// TODO: add comments

import (
	"fmt"
)

type Point struct {
	X, Y int
}

type Rect0 struct {
	Min, Max Point
}

func (r *Rect0) pointRange(min_x, min_y, max_x, max_y int) {
	r.Min = Point{min_x, min_y}
	r.Max = Point{max_x, max_y}
}

func main() {

	arr0 := make([]int, 2, 5)

	arr1 := new([]int)

	rect0 := new(Rect0)

	fmt.Printf("\ntype: %T | len (index range): %v | cap (slice operation range): %v | values: %v\n\n", arr0, len(arr0), cap(arr0), arr0)

	fmt.Printf("type: %T | len (index range): %v | cap (slice operation range): %v | values: %v\n\n", arr1, len(*arr1), cap(*arr1), arr1)

	fmt.Printf("Rect0 allocated with the new function: %+v\n\n", rect0)

	rect0.pointRange(10, 20, 50, 60)

	fmt.Printf("Rect0 values: %+v\n\n", rect0)

	fmt.Printf("slice: %v\n\n", arr0[1])

}
