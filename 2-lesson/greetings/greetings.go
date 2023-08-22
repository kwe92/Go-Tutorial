package greetings

import "fmt"

func Hello(name string) (msg string) {
	message := fmt.Sprintf("Hello %v. Welcome to the Go Developer Community", name)
	return message
}

// TODO: Maybe transfer to written notes

// Summary

//   ~ Created a package called greetings

//   ~ Imported the fmt package

//       - used to format and print text to the terminal

//   ~ Declared a function Hello

//       - the Hello function has one parameter 'name'
//         with an explictly defined type of `string`
//         implying that 'name' takes as an argument a string value
//         Hello then returns a string as a value
//         as indicated from the second head which specifies return type

//   ~ Function Names

//       + Private Functions

//           - snake-cased function are private to the package

//       + Go Exported Functions

//           - functions that start with a capital letter are
//             go exported functions and can be used in other packages

//   ~ Declaring and Initializing Variables

//       + The := Operator vs = Operator

//           - in Go there is an added level of granularity
//             with declaring and initializing variables with operators
//           - you can just declare a variable with the `=` operator
//             and initalize it later
//           - or you can declare and initialize a variable
//           - simultaneously with the := operator

//   ~ fmt.Sprintf | returns string

//       - fmt.Sprintf takes a string and a format specifier and
//         formats the passed in string according to the format specifier
//       - the string is then returned
//       - slightly confusing by its name but it does not
//         actually print anything to the terminal
