// Day 3

package main

import (
	"fmt"
)

func main() {

	var n uint

	fmt.Scanf("%d", &n)

	if n%2 == 0 && n >= 2 && n <= 5 {
		fmt.Println("Not Weird")
	} else if n%2 == 0 && n >= 6 && n <= 20 {
		fmt.Println("Weird")
	} else if n%2 == 0 && n > 20 {
		fmt.Println("Not Weird")
	} else if n%2 == 1 {
		fmt.Println("Weird")
	} else {
		fmt.Println("Not Weird")
	}
}
