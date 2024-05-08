package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {

	var counter int

	var changeCounter float64

	var changeAmounts []float64

	// var remainders []int

	for counter < 100 {
		counter += 1
		changeCounter += .01

		changeAmounts = append(changeAmounts, roundToNearest(changeCounter))
	}

	fmt.Println(changeAmounts)

	for _, change := range changeAmounts {
		fmt.Println(int(change * 100))

	}

	// for _, val := range change {
	// 	fmt.Println(getRemanider())

	// }

	// transactionAmount := 67

	// remainder, err := getRemanider(transactionAmount)

	// if err != nil {
	// 	return
	// }

	// fmt.Printf("remainder: %.2f\n", (float64(remainder) / 100))

}

// GetRemanider: expects an integer representation of currency and returns the remainder that is needed to round to the nearest dollar.
//
// e.g. 12550 == $125.50 returning the integer 50 to get 12550 + 50 == $125.50 + $0.50 == $126.00
func getRemanider(transactionAmount int) (int, error) {

	// convert to float
	transactionAmountInUSD := float64(transactionAmount) / 100

	// convert float to a string with trailing 0's if there is no change
	stringTransactionAmountInUSD := fmt.Sprintf("%.2f", transactionAmountInUSD)

	fmt.Println(stringTransactionAmountInUSD)

	// split string on `.` and get change which is the second element
	stringChange := strings.Split(fmt.Sprint(stringTransactionAmountInUSD), ".")[1]

	// convert string remainder change to integer
	change, err := strconv.Atoi(stringChange)

	fmt.Println(change)

	if err != nil {
		log.Printf("ERROR in GetRemanider: %s", err.Error())

		return 0, err
	}

	if change == 0 {
		return 100, nil
	}

	remainder := 100 - change

	return remainder, nil

}

func roundToNearest(val float64) float64 {
	return math.Round(val*100) / 100
}

//------------------------------------ Original ------------------------------------//

// func main() {

// 	var counter float64

// 	var change []float64

// 	for counter < 1000 {
// 		counter += 10.01
// 		change = append(change, roundToNearest(counter))
// 	}

// 	fmt.Println(change)

// 	for _, val := range change {
// 		fmt.Println(val + getRemainder(val))

// 	}

// }

// func getRemainder(initialTransaction float64) float64 {
// 	remainder := 1 - (initialTransaction - math.Floor(initialTransaction))

// 	remainder = roundToNearest(remainder)

// 	if remainder == 0 {
// 		return 1
// 	}

// 	return remainder
// }
