package main

import "fmt"

type Person struct {
	firstName  string
	middleName *string
	lastName   string
}

func main() {
	person := Person{
		firstName:  "Kweayon",
		middleName: returnPointer[string]("LaCharles"),
		lastName:   "Clark",
	}

	fmt.Println(person)
}

// returnPointer: a generic function that returns a pointer to a value of any type.
func returnPointer[T any](value T) *T {
	return &value
}

// struct: Pointer Types For Fields

//   - if a field has a type of pointer type then the assigned value must be a pointer type
//   - fields with pointer types may not be adjacent to eachother in memory

// Primitive Literal Pointers

//   - primitives can not be prefixed with & address operator
//   - you must assign the value to a variable and point to that variable or return a pointer from a function
