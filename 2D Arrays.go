// Day 11

package main

import "fmt"

func main() {
	var m [6][6]int
	var n int
	hasil := 0
	max := 0

	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			fmt.Scan(&n)
			m[i][j] = n
		}
	}

	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			if i+2 < 6 && j+2 < 6 {
				hasil = m[i][j] + m[i][j+1] + m[i][j+2] + m[i+1][j+1] + m[i+2][j] + m[i+2][j+1] + m[i+2][j+2]
				if hasil >= max {
					max = hasil
				}

			}
		}
	}

	fmt.Println(max)
}
