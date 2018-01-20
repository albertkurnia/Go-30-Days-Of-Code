// Day 5

package main

import (
	"fmt"
)

func main() {
	var T int

	fmt.Scan(&T)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", T, i, T*i)
	}
}
