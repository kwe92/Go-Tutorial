package main

import "fmt"

type Person struct {
	Fname string
	Lname string
	Age   int
}

func main() {

	// declare a variable of the user defined type Person
	var animeCharacter0 Person

	// initialize with struct literal and named fields
	animeCharacter0 = Person{Fname: "Naruto", Lname: "Uzumaki", Age: 14}

	fmt.Println(animeCharacter0)

	// declare and initialize with struct literal and unnamed fields
	animeCharacter1 := Person{"Shikamaru", "Nara", 14}

	fmt.Println(animeCharacter1)

}

// struct GO

//   - a data structure that's comprised of a set of fields
//   - can implement methods on structures which will be covered later

// Variables Declared as struct

//   - variables that are declared as a struct without initialization have their fields
//     initialzed to their zero value

// Struct Fields In Memory

//   - struct fields that are not pointer types are adjacent to eachother in memory
