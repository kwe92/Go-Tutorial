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

//   - panics are generated when the GO runtime is unable to determine what to do next
//   - panics can also be generated when there are environmental issues like
//     attempting to read past the end of a slice or memory runs out
//   - exits the program with a message and a stack trace

// defer after panic

//   - when the GO runtime panics the current function exits and all deferred functions run
//     in last in first out order until the main function is reached
//   - the program then exits with a message and a stack trace

// recover

//   - provides a way to capture a panic and gracefully shutdown
//     or continue execution dispit a panic occuring
//   - additional use cases:
//       ~ log panic before the program exits
//   	 ~ convert panics into errors so the users of an Api can handle the error

// deferred recovery

//   - once a panic occurs deferred functions are executed
//   - implying that a recovery must be deferred in order to capture a panic

// panic and recover | NOT TYPICAL Exception Handling

//   - panic should only be used in fatal situations

// Continuing Execution After panics

//   - continuing execution after a panic should be avoided as it can cause more issues
//   - recovery should be used to log a panic and stop execution with os.Exit(1) or convert a panic to an error
//   - recover isn't very verbose in its description on what failed which is not very idiomatic in GO

// Panics in Third Party Libraries You Create

//   - a panic should never escape your API
//   - if there are posibilities of a panic in an API you create use panic and recovery
//     to convert the panic into an error and return the error to the user of your API

// You Want a Stack Trace?

//   - don't panic, use a third party library or create your own stack trace with error wrapping

// Stack Traces and Paths

//   - stack traces expose the full path to the file on the computer where the program was compiled
//   - use the -trimpath flag to replace the full path with just the package when building your binary
