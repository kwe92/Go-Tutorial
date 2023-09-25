package main

import (
	"golang.org/x/tour/pic"
	//"fmt"
)

func Pic(dx, dy int) [][]uint8 {
	arr := make([][]uint8, dy)
	for y := range arr {
		arr[y] = make([]uint8, dx)
		for x := range arr[y] {
			arr[y][x] = uint8((x + y) * (x * y / 4) / (x ^ y + 1))
		}
	}
	return arr
}

func main() {
	// arr := make([][]uint8, 5)
	// fmt.Println(arr)

	pic.Show(Pic)
}
