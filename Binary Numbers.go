// Day 10

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var N int64
	var max int = 0
	var count int = 0

	fmt.Scan(&N)

	m := strconv.FormatInt(N, 2)
	for _, k := range m {
		if k == 49 {
			count++
			if count >= max {
				max = count
			}
		} else if k == 48 {
			count = 0
		}
	}

	fmt.Println(max)
}
