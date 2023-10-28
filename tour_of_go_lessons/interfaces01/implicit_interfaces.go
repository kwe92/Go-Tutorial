package main

import (
	"fmt"
	"strconv"
)

type Stringer interface {
	String() string
}

func ToString(any interface{}) string {

	if v, ok := any.(Stringer); ok {

		return v.String()

	}
	switch v := any.(type) {

	case int:
		return strconv.Itoa(v)

	case float64:
		return strconv.FormatFloat(v, 'g', 0, 64)

	}

	return "???"
}

type Binary uint64

func (i Binary) String() string {
	return strconv.FormatUint(i.Get(), 2)
}

func (i Binary) Get() uint64 {
	return uint64(i)
}

var binary Binary

func main() {
	var s string = "Trust in the Lord with all your heart."

	var integer int = 42

	var float float64 = 3.14

	binary = Binary(integer)

	var arr = []interface{}{binary, s, integer, float}

	for _, ele := range arr {
		val := ToString(ele)

		fmt.Println(val)
	}
}

// Interfaces Are Implemented Implicitly

//     - types implement interfaces by implementing the associated methods `method set` of an interface
//     - GO removes the declaration of intent like the implements keyword that you may find in other object oriented languages
//     - If a type implements an interfaces method signatures `method set` then it can be considered an implementation of that interface implicitly

// Type Conversion

//   - type conversion can be done with both primative and defined types
//   - thats why the invocation of Binary(integer) works

// Explicit Interface-to-Interface runtime Checks With Type Assertions

//   - type assertions allow you to check the dynamic type of a value at runtime

// strconv Package

//   - A package used for string conversion

// Explination of Above Code

//   1. defined a type Stringer
//   2. implemented a ToString method
//     2.a the if statement does a type assertion checking if the value is of type Stringer
//     2.b if the predicate returns true then the string method is invoked on that value
//   3. if the if statements predicate returns false then a switch statement does a type assertion against with various cases
//   4. the type Binary is defined as a uint64 `unsigned 64bit integer`
//   5. the methods String and Get are implemented for the type Binary
