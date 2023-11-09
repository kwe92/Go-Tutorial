package main

import (
	"fmt"
	"time"
)

func say(something string) {

	for i := 0; i < 5; i++ {
		// wait 100 milliseconds before executing the print statement
		time.Sleep(100 * time.Millisecond)
		fmt.Println(something)

	}
}

func main() {
	go func() {
		say("Begin, to begin is half the work.")
	}()
	say("Let half still remain.")
}

// Concurency

//   - multiple computations happening simultaneously
//   - breaking a process into individual components `threads`

// Concurrency Definition by Google io

//   - the composition of independently executing computations
//   - can be viewed as a model for software construction in which software is constructed in a way
//     where clean code is written and interacts well with the real world

// Process

//   - an instance of a program that was given partitioned resources by the operating system during the virtualization process
//   - processes run in isolation and do not have access to another processes resources
//   - processes are composed of one or more threads

// Threads

//   - the most granular set or sequence of instructions a computer can manage or execute
//   - the operating system gives these units of execution time to run
//   - threads that are within the same process share resources of that process

// CPU and Threads

//   - the CPU can execute instructions from one or more threads within a single process simultaneously

// Operating System and Threads

//   - the operating system has a process scheduler that schedules threads `units of execution` on the CPU
//     allowing every process and thread within a process to run

// goroutines

//   - independently executing functions launched by the go statement
//   - lightweight threads managed by the GO runtime
//   - they can be seen as extremely cheap threads
//   - goroutines do not make the caller wait
//   - goroutines have their own call stack and they are elastic `can grow and shrink`
//   - a single threaded program can have multiple goroutines
//?  - goroutines are multiplexed onto threads as needed

// Launching a Function as a goroutine

//   - there are no async or await keywords like other languages that handle asnychronous operations
//   - to indicate a function will run a concurrent process the `go` keyword preceeds a function invocation
//   - any function can be prefixed with the go keyword but it is idiomatic to wrap the function in a closure

// Main Function: The Initial goroutine

//   - when a Go program starts the Go Runtime launches a single goroutine and creates a number of threads
//   - the Go Runtime Scheduler automatically assigns the initial goroutine `main function` and all subsequent
//     goroutines to the threads created by the Go Runtime
//   - The Go Runtime Scheduler is similar to the operating system process scheduler but cheaper
//     as you are not creating an operating system level resource

// goroutine Address Space

//   - goroutines run in the same address space
//   - Memory access that is shared must be synchronized
//   - to synchronize memory you can use primitives in GO or the sync package

// ?--- What is the GO Runtime ---?
