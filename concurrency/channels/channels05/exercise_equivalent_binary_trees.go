package main

// TODO: figure out why your goroutine is deadlocking

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func main() {

	btree := tree.New(1)
	ch := make(chan int)

	go inOrderTraversal(btree, ch)

	for ele := range ch {
		fmt.Println("from main: ", ele)
	}

	// fmt.Println(btree.Value)

}

func inOrderTraversal(btree *tree.Tree, ch chan int) {
	if btree != nil {
		inOrderTraversal(btree.Left, ch)
		// fmt.Printf("%d\n", btree.Value)
		ch <- btree.Value
		inOrderTraversal(btree.Right, ch)

	}
}

// func inOrderTraversal(btree *tree.Tree) {
// 	if btree != nil {
// 		inOrderTraversal(btree.Left)
// 		fmt.Printf("%d\n", btree.Value)
// 		inOrderTraversal(btree.Right)

// 	}
// }
