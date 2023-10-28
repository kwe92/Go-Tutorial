package main

import "fmt"

func main() {

	var person00 struct {
		Fname string
		Lname string
		Age   int
	}

	person00.Age = 14
	person00.Fname = "Gaara"
	person00.Lname = "Sabaku"

	fmt.Println(person00)

	person01 := struct {
		Fname string
		Lname string
		Age   int
	}{
		"Sasori", "", 35,
	}

	fmt.Println(person01)

	var person02 = struct {
		Fname string
		Lname string
		Age   int
	}{
		Fname: "Obito",
		Lname: "Uchiha",
		Age:   31,
	}

	fmt.Println(person02)

}

// Anonymous struct's

//   - structs in go can be anonymous `not concretly defined`
