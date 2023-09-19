package main

import "fmt"

type Emplyee struct {
	id     int
	fname  string
	lname  string
	salary float32
}

func spacedPrint(str string) {
	fmt.Println("\n" + str + "\n")
}

func main() {
	intArr := [3]int{1, 2, 3}

	const variable0 = 4

	fmt.Printf("\n%v\n", variable0)

	emp1 := Emplyee{id: 1101, fname: "Kweayon", lname: "Clark", salary: 130000}

	spacedPrint("Hello Universe.")
	fmt.Printf("%v\n\n", intArr)
	fmt.Printf("%v\n\n", emp1)

}

// fmt
//   - a package part of the standard library
//   - used to format and print strings
