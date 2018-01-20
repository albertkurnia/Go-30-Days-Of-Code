// Day 6

package main

import (
	"fmt"
	"strings"
)

func main() {
	var T int
	var str string
	var arr []string

	fmt.Scanln(&T)

	for i := 0; i < T; i++ {
		fmt.Scanln(&str)
		arr = append(arr, str)
	}

	for _, v := range arr {
		strs := strings.Split(v, "")
		odd := ""
		even := ""
		for i, k := range strs {
			if i%2 == 0 {
				odd += k
			} else {
				even += k
			}
		}

		fmt.Println(odd + " " + even)
	}
}
