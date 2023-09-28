package main

import "fmt"

func main() {

	var value MyValue = MyValue{5}

	any := (interface{})(5)

	fmt.Println("\nValue: ", value)

	Adder(&value, 35)

	printType(value)

	printType(any)

}

func printType(any interface{}) {
	fmt.Printf("\nvalue: %v | type %T\n", any, any)
}

func Adder(value Settable, add int) {

	value.SetValue(add)

}

type Settable interface {
	SetValue(interface{})
}

type MyValue struct {
	Val int
}

func (m *MyValue) SetValue(val interface{}) {

	m.Val = val.(int)
}

func (m *MyValue) Abs() int {
	if m.Val < 0 {
		return m.Get()
	}
	return m.Get()
}

func (m *MyValue) Get() int {
	return int(m.Val)
}

// Interface Values

//   - An interfaces value can be viewed as a tuple of pointers or a two word pointer pair in memory
//     comprised of the value and the values concrete type --> (value, data_type)
//   - the pointer to the data type becomes a list of function pointers
//   - the value is a pointer to a copy of the original value? Check this for validation

// concrete type

// - The type that implements the associated interfaces method signatures

// Go Method Tables / i-Table

//   - computed at run-time
