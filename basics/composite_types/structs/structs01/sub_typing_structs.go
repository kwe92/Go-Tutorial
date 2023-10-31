package main

import "fmt"

type Person struct {
	Fname string
	Lname string
	Age   int
}

func (p *Person) String() string {
	return fmt.Sprintf("Person{Fname: %s, Lname: %s, Age: %d}", p.Fname, p.Lname, p.Age)
}

type Human Person

func main() {
	var animeCharacter00 Person
	var animeCharacter01 Human

	// results in a compile time error
	// fmt.Println(animeCharacter00 == animeCharacter01)

	fmt.Println(animeCharacter00)

	fmt.Println(animeCharacter01)

	// anonymous struct
	var animeCharacter02 = struct {
		Fname string
		Lname string
		Age   int
	}{
		Fname: "Shikamaru",
		Lname: "Nara",
		Age:   17,
	}

	fmt.Println(animeCharacter02)

	animeCharacter00 = animeCharacter02

	animeCharacter01 = animeCharacter02

	fmt.Println(animeCharacter00.String())

	fmt.Println(animeCharacter01)

	// results in a compile time error
	// Human type does not implement a String method so there is not one to call
	// fmt.Println(animeCharacter01.String())

	// results in a compile time error
	// animeCharacter00 = animeCharacter01

}

// Sub-Typing is not Iheritance

//   - sub-types have the same field names, field order, and field types as the parent type
//   - sub-types do not share `inherit` any methods defined in the parent type
