// TODO: add more operations to the linked list
// TODO: create a better implementation

package main

import (
	"fmt"
	"reflect"
)

type Node[T any] struct {
	next *Node[T]
	val  T
}

func insertAtBeginning[T any](node Node[T], val T) Node[T] {

	newNode := Node[T]{val: val}

	if reflect.TypeOf(node.val).Kind() == reflect.Int {
		if any(node.val).(int) == 0 || any(node.val) == nil {
			return newNode
		}
	}

	newNode.next = &node

	return newNode

}

func main() {
	linkedList := Node[int]{}

	newList0 := insertAtBeginning(linkedList, 20)

	newList1 := insertAtBeginning(newList0, 42)

	fmt.Printf("\nlinkedList: %v", linkedList)

	fmt.Printf("\nnewList0: %v", newList0)

	fmt.Printf("\nnewList1: %v", newList1)

	fmt.Printf("\nnewList1 value: %v\n\n", newList1.next.val)

}

// Generic Types

//   - defined types can also have type patameters
//   - allowing the implementation of generic data structures
