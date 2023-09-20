package main

// TODO: use a diffrent example

import (
	"fmt"
	"strings"
)

type I interface {
	Shout()
}

func main() {

	var i I = T{"take therefore no thought for the morrow; For the morrow shall take thought for the things of itself. sufficient unto the day is the evil thereof."}

	i.Shout()
}

type T struct {
	S string
}

func (t T) Shout() {
	fmt.Println(strings.ToUpper(t.S))
}

// Interfaces Are Implemented Implicitly

//     - types implement interfaces by implementing the associated methods of an interface
//     - GO removes the declaration of intent like the implements keyword that you may find in other object oriented languages
//     - If a type implement’s an interfaces method signatures then it can be considered an implementation of that interface implicitly
