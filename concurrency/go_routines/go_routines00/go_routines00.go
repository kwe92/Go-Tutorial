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

//   - multiple computations happening at the same time
//   - a process is broken into individual components

// Process

//   - an instance of a program that was given partitioned resources by the operating system during the virtualization process
//   - processes run in isolation and do not have access to another processes resources
//   - a processes are composed of one or more threads

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

//   - lightweight processes managed by the GO runtime
//   - they can be seen as extremly cheap threads

// Launching a Function as a goroutine

//   - there are no async or await keywords like other languages that handle asnychronous operations
//   - to indicate a function will run a concurrent process the `go` keyword preceeds a function invocation
//   - any function can be prefixed with the go keyword but it is idiomatic to wrap the function in a closure

// When a Go Program Starts | change the name later

//   - when a Go program starts the Go Runtime launches a single goroutine and creates a number of threads
//   - the initial goroutine and all other goroutines are automatically assigned to the threads
//     that the Go Runtime created by the Go Runtime Scheduler `similar to the operating system process scheduler but cheaper`

// goroutine Address Space

//   - Go routines run in the same address space
//   - Memory access that is shared must be synchronized
//   - to synchronize memory you can use primitives in GO or the sync package

// ?--- What is the GO Runtime ---?
