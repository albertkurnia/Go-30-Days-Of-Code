// Day 2

package main

import (
	"fmt"
)

func main() {
	var mealCost float64
	var tipPercent uint
	var taxPercent uint

	fmt.Scanf("%f\n %d\n %d\n", &mealCost, &tipPercent, &taxPercent)

	tax := mealCost * float64(taxPercent) / 100
	tip := mealCost * float64(tipPercent) / 100

	fmt.Printf("The total meal cost is %.1f dollars.", mealCost+tax+tip)
}
