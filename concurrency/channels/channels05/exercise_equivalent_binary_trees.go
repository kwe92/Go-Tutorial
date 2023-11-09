package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func main() {

	btree0 := tree.New(1)

	btree1 := tree.New(2)

	fmt.Println(Same(btree0, btree1))

}

// Walk: invokes inOrderTraversal to walk the binary tree and then close the channel
func Walk(btree *tree.Tree, ch chan int) {
	inOrderTraversal(btree, ch)
	defer close(ch)
}

// inOrderTraversal: traverse a binary tree starting from the left subtree
func inOrderTraversal(btree *tree.Tree, ch chan int) {
	if btree != nil {
		inOrderTraversal(btree.Left, ch)
		// write binary tree values to the channel
		ch <- btree.Value
		inOrderTraversal(btree.Right, ch)
	}
}

// Same: compare the values of two binary trees
func Same(t0, t1 *tree.Tree) bool {
	ch0, ch1 := make(chan int), make(chan int)

	go Walk(t0, ch0)
	go Walk(t1, ch1)

	for {
		// read from channels and check if they are open
		v0, ok0 := <-ch0
		v1, ok1 := <-ch1

		// check if both channels are open and if both channels have the same values
		if ok0 != ok1 || v0 != v1 {
			return false
		}

		// indicates the channel is closed breaking out of the for loop
		if !ok0 {
			break
		}

	}
	return true

}
