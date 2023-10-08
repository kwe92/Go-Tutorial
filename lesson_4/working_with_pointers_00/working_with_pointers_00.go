package main

import (
	"fmt"
)

type TaskRecord struct {
	ID          string
	TaskName    string
	TaskDetails string
}

var tasks []TaskRecord

// addTask takes a pointer to an array of TaskRecord's and a TaskRecord as arguments and appends the task to the array, modifying the task array passed in.
func addTask(taskArr *[]TaskRecord, task TaskRecord) {

	// point to the array and append values to it
	*taskArr = append(*taskArr, task)

}

// addTaskNoPointer takes an array of TaskRecord's and a TaskRecord as arguments and appends the task to a copy of the task array, not modifying the passed in array.
func addTaskNoPointer(taskArr []TaskRecord, task TaskRecord) {

	taskArr = append(taskArr, task)

}

func main() {
	task0 := TaskRecord{
		ID: "1001", TaskName: "Begin", TaskDetails: "Begin, to begin is half the work.",
	}

	task1 := TaskRecord{ID: "1002", TaskName: "Zazen", TaskDetails: "sitting is enough."}

	addTask(&tasks, task0)

	addTaskNoPointer(tasks, task1)

	fmt.Println(tasks)
}

// Memory

//   - a storage system for many diffrent types of data
//   - Can be viewed as an indexable collection that stores objects for a program
//   - Objects in memory are indexable by memory addresses, offering constant time operations O(1)
//        ~ Insert, Search, Update, Delete

// Pointers

//   - Pointers point to the value of an object in memory by the location in memory `memory address`
//     where the object is stored
//   - Reducing memory allocation (memory storage is finite)

// Ampersand Operation

//   - &Object generates a pointer to the operand (suffixed object)

// Pointer Value

//   - the asterisk operation denotes a pointers underlying value
//   - used to update the value at the point in memory

//? Pointers In Practice

// Functions Modifying Collections

//   - Collections passed as arguments to functions are copied within the function
//   - which means that the parameter used within the function will not be the original object
//   - implying that any changes to the collection passed in will not be reflected outside of the function

// Modify by Pointer

//   - to modify an object passed into a function as an argument
//     you must pass in a pointer to that object
//   - Which can then be used to modify the object at that point in memory
