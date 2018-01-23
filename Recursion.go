// Day 9

package main

import "fmt"

func factorial(N int) int {
	if N == 1 {
		return 1
	}
	return factorial(N-1) * N
}

func main() {

	var N int

	fmt.Scan(&N)

	fmt.Println(factorial(N))
}
