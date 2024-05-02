package main

import (
	"fmt"
	"math"
)

func main() {

	var counter float64

	var change []float64

	for counter < 1000 {
		counter += 10.01
		change = append(change, roundToNearest(counter))
	}

	fmt.Println(change)

	for _, val := range change {
		fmt.Println(val + getRemainder(val))

	}

}

func getRemainder(initialTransaction float64) float64 {
	remainder := 1 - (initialTransaction - math.Floor(initialTransaction))

	remainder = roundToNearest(remainder)

	if remainder == 0 {
		return 1
	}

	return remainder
}

func roundToNearest(val float64) float64 {
	return math.Round(val*100) / 100
}
