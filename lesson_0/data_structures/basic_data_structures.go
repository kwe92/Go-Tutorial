package main

import "fmt"

// TODO: add comments

func main() {
	i := 1234
	j := int32(1)
	f := float32(3.14)

	bytes := []byte{'h', 'e', 'l', 'l', 'o'}

	primes := []int{2, 3, 5, 7}

	arr0 := [6]any{i, j, f, bytes, string(bytes), primes}

	fmt.Println(arr0)

}
