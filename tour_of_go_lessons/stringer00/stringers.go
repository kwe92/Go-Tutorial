package main

import "fmt"

func main() {

	person0 := Person1{Fname: "David", Lname: "The King"}

	person1 := Person0{Fname: "John", Lname: "The Baptist"}

	fmt.Println("\ndefault struct string representation:\n\n", person0)

	fmt.Println("\noverridden struct string representation:\n\n", person1)

}

type Stringer interface {
	String() string
}
type Person0 struct {
	Fname string
	Lname string
}

type Person1 struct {
	Fname string
	Lname string
}

func (p Person0) String() string {
	return fmt.Sprintf("first name: %v | last name: %v", p.Fname, p.Lname)
}

// Stringer interface

//   - a type that can describe itself as a string
//   - found often throughout GO
//   - fmt package and many others look for the Stringer interface to print values
//   - similar to __repr__ and __str__ methods in python classes
//     or overriding toString method in Dart classes

// Process to Print Custom String Representations

// 1. define a type
// 2. implement a String method with a value receiver that returns a string
//    representation of that type
