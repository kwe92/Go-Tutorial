package main

// TODO: create better example

import (
	"fmt"
	"sync"
)

type ProtectedData struct {
	data int
	lock *sync.Mutex
}

func main() {

	var waitGroup = &sync.WaitGroup{}

	protectedData := &ProtectedData{
		data: 0,
		lock: &sync.Mutex{},
	}

	fmt.Println(protectedData)

	n := 20

	waitGroup.Add(n)

	for i := 0; i < n; i++ {
		go addTwo(protectedData, waitGroup)

	}
	waitGroup.Wait()

	fmt.Println(protectedData)

}

func addTwo(data *ProtectedData, waitGroup *sync.WaitGroup) {
	data.lock.Lock()

	data.data += 2

	data.lock.Unlock()

	waitGroup.Done()
}
