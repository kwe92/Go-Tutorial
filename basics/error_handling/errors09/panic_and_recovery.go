package main

import (
	"fmt"
)

func doPanic(msg string) {
	panic(msg)
}

func funcThatPanics() {
	defer func() {
		if paniced := recover(); paniced != nil {
			fmt.Println("THERE WAS A PANIC: ", paniced)

		}
	}()
	doPanic("MONSTER!")
}

func main() {
	// fmt.Println(os.Args[1:][0])

	funcThatPanics()

	fmt.Println("What kind though? It may be petable!")

}

// panic

//   - generated when the GO runtime is unable to figureout what to do next
//   - panics can also be generated when there are environmental issues like
//     attempting to read past the end of a slice or memory runs out
//   - exits the program with a message and a stack trace

// defer after panic

//   - when the GO runtime panics the current function exits and all defered functions run
//     in last in first out order until the main function is reached
//   - the program then exits with a message and a stack trace

// recover

//   - provides a way to capture a panic and gracefully shutdown or continue execution dispit a panic occuring

// defered recovery

//   - once a panic occurs only differed functions can run
//   - this implies that a recovery must be defered in order capture a panic

// panic and recover is NOT TYPICAL Exception Handling

//   - panic should only be used in fatal situations

// Continuing Execution After panics

//   - typically you want the program to exit as the continuation of execution after a panic can cause more issues
//   - you can do some preprocessing like logging panics durring recovery and exit execution with os.Exit(1)
//   - recover isn't very verbose in its description on what failed which is not very idiomatic in GO

// Panics in Third Party Libraries You Create

//   - a panic should never escape your API
//   - if there is a posibility of a panic in an API you create
//   - use panic and recovery to convert the panic into an error and return the error to the user of your API

// You Want a Stack Trace?

//   - don't panic, use a third party library or create your own stack trace with error wrapping

// Stack Traces and Paths

//   - stack traces expose the full path to the file on the computer where the program was compiled
//   - use the -trimpath flag to replace the full path with just the package when building your binary
