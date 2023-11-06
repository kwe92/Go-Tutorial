package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter: safe to use concurrently
type SafeCounter struct {
	mu  sync.Mutex
	val map[string]int
}

func (sc *SafeCounter) Increment(key string) {
	// invoke lock to protect the value
	sc.mu.Lock()

	sc.val[key]++

	// unlock protected value after doing something
	sc.mu.Unlock()
}

func (sc *SafeCounter) Value(key string) int {

	sc.mu.Lock()

	// we defer because we are returning a value and want to unlock after the value is returned
	defer sc.mu.Unlock()

	return sc.val[key]
}

func main() {
	safeCounter := SafeCounter{
		val: map[string]int{},
	}

	for i := 0; i < 10; i++ {
		go safeCounter.Increment("steps")
		go safeCounter.Increment("steps")

	}

	time.Sleep(time.Second)
	fmt.Println(safeCounter.Value("steps"))
}

// Mutex

//   - mutex is a data structure that implements the concept of mutual exclusion

// Mutual Exclusion

//   - prevents multiple threads from accessing the same object simultaneously
//   - one thread has exclusive access to a shared resource until it is unlocked
//   - limits the concurrent execution of some code or access to a shared piece of data
//   - goroutines can only access a mutually exclusive object one at a time
//   - used to protect a value in a concurrent process across threads
//   - only one thread can Lock at any given time but any thread can Unlock

// Critical Section

//   - the protected part of execution or shared data

// sync.Mutex

//   - a struct part of the sync package that implements mutual exclusion
//   - defines two methods Lock and Unlock

// Mutex and Data Obscurity

//   - mutex's do not indicate which go-routine has access to the value
//   - the value is shared between all concurrent processes
//   - the order of processing is hard to understand

// When to Use Mutual Exclusion

//   - when go-routines read or write a shared value but the value is not processed
