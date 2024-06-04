package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {

	// incrementer
	var counter int

	// incrementer for change amounts
	var changeCounter float64

	// array of change amounts ranging from $0.01 to $1.00
	var changeAmounts []float64

	// interger representation of nearest amount of change required to round to nearest dollar
	var nearestDollarAmounts []int

	for counter < 100 {

		counter += 1

		changeCounter += .01

		changeAmounts = append(changeAmounts, roundToNearest(changeCounter))
	}

	fmt.Println(changeAmounts)

	for _, change := range changeAmounts {

		nearestAmount, _ := getNearestAmount(int(change * 100))

		nearestDollarAmounts = append(nearestDollarAmounts, nearestAmount)

	}

	fmt.Println(nearestDollarAmounts)

	for i := 0; i < len(changeAmounts); i++ {

		fmt.Printf("%d + %d  = %d\n", int(changeAmounts[i]*100), nearestDollarAmounts[i], (int(changeAmounts[i]*100) + nearestDollarAmounts[i]))

	}

}

// getNearestAmount: expects an integer representation of currency and returns the remainder that is needed to round to the nearest dollar.
//
// e.g. 12550 == $125.50 returning the integer 50 to get 12550 + 50 == $125.50 + $0.50 == $126.00
func getNearestAmount(transactionAmount int) (int, error) {

	// convert to float
	transactionAmountInUSD := float64(transactionAmount) / 100

	// convert float to a string with trailing 0's if there is no change
	stringTransactionAmountInUSD := fmt.Sprintf("%.2f", transactionAmountInUSD)

	// split string on `.` and get change which is the second element
	stringChange := strings.Split(fmt.Sprint(stringTransactionAmountInUSD), ".")[1]

	// convert string remainder change to integer
	change, err := strconv.Atoi(stringChange)

	if err != nil {
		log.Printf("ERROR in GetRemanider: %s", err.Error())

		return 0, err
	}

	if change == 0 {
		return 100, nil
	}

	nearestAmount := 100 - change

	return nearestAmount, nil

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
