package main

import "fmt"

type MailCategory int

const (
	Uncategorized MailCategory = iota
	Personal
	Spam
	Social
	Advertisements
)

func main() {
	fmt.Println(Social)
}

// GO Enumerations

//   - GO does not have an enumeration type
//   - GO has the concept of iota

// iota as Enumerations

//   - iota allows you to assign increasing values to a set of constants
//   - useful if you dont need the value
