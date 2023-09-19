package main

import "fmt"

// nil map, can't be assigned keys or values
var m map[string]interface{}

func main() {

	// allocate memory and initialize map
	m = make(map[string]interface{})

	m["Answer"] = 42

	fmt.Println("\nThe Answer to The Universe: ", m["Answer"])

	updateIfExists(&m, "Answer", 48)

	fmt.Println("\n\nThe Answer to The Universe: ", m["Answer"])

	updateIfExists(&m, "I_Dont_Exist", 42)

	deleteIfExists(&m, "Answer")

	fmt.Printf("\n\nThe Answer to The Universe: %v\n\n", m["Answer"])

}

// updateIfExists updates a value at a specified key if it exists.
func updateIfExists(m *map[string]interface{}, key string, value interface{}) bool {

	var isPresent bool

	previousValue, isPresent := (*m)[key]

	if isPresent {
		fmt.Printf("\nupdate from: %v", previousValue)
		(*m)[key] = value
		fmt.Printf("\n\nupdate to: %v", (*m)[key])

		return isPresent
	}
	fmt.Printf("\nThe key: %v is not present in the hashMap", key)
	return isPresent
}

// deleteIfExists updates a value at a specified key if it exists.
func deleteIfExists(m *map[string]interface{}, key string) bool {

	var isPresent bool

	_, isPresent = (*m)[key]

	if isPresent {
		delete((*m), key)
		fmt.Printf("\n\nDeleted value at key: %v", key)
		return isPresent
	}
	fmt.Printf("\nThe key: %v is not present in the hashMap", key)
	return isPresent
}
