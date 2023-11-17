package main

import (
	"fmt"
	"sync"
)

// SafeCounter: protected data safe for concurrent use
type SafeCounter struct {
	lock *sync.Mutex
	val  map[string]int
}

func (sc *SafeCounter) Increment(key string) {

	// lock data structure to protect value
	sc.lock.Lock()

	sc.val[key]++

	// unlock data structure to unprotect value after doing something
	sc.lock.Unlock()
}

func (sc *SafeCounter) Value(key string) int {

	sc.lock.Lock()

	// defer Unlock because we are returning and want to unlock after returning
	defer sc.lock.Unlock()

	return sc.val[key]
}

func main() {

	safeCounter := SafeCounter{
		val:  map[string]int{},
		lock: &sync.Mutex{},
	}

	n := 10

	waitGroup := &sync.WaitGroup{}

	waitGroup.Add(n * 2)

	for i := 0; i < n; i++ {

		go func() {
			safeCounter.Increment("steps")
			waitGroup.Done()
		}()

		go func() {
			safeCounter.Increment("steps")
			waitGroup.Done()
		}()

	}

	waitGroup.Wait()

	fmt.Println(safeCounter.Value("steps"))
}

// Mutex

//   - a shared memory locking mechanism used synchronize the access of a shared mutable resource
//     between multiple concurrent operations
//   - a data structure that implements the concept of mutual exclusion

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
