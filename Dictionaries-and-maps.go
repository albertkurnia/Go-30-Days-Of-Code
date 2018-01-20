// Day 8

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N int
	scanner := bufio.NewScanner(os.Stdin)

	m := make(map[string]int)

	fmt.Scanln(&N)

	for i := 0; i < N; i++ {
		var str string
		var num int

		fmt.Scan(&str)
		fmt.Scanln(&num)

		m[str] = num
	}

	for scanner.Scan() {
		var str string

		fmt.Scanln(&str)

		fmt.Printf("%s=%d\n", str, m[str])
	}
}
