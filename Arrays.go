// Day 7

package main

import "fmt"

func main() {

	var N int

	fmt.Scanln(&N)
	var arr [1001]int

	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	for i := N - 1; i >= 0; i-- {
		fmt.Printf("%d ", arr[i])
	}
}
