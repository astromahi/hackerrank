// https://www.hackerrank.com/contests/30-days-of-code/challenges/day-11-more-review
// Day 11: 2D-Arrays + More Review!

package main

import (
	"fmt"
)

func main() {
	var n, sum, largest int
	var matrix [6][6]int

	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			fmt.Scanf("%d", &n)
			matrix[i][j] = n
		}
	}

	sum, largest = 0, -9223372036854775808
	for k := 0; k <= 3; k++ {
		for l := 0; l <= 3; l++ {
			sum = matrix[k][l] + matrix[k][l+1] + matrix[k][l+2] + matrix[k+1][l+1] + matrix[k+2][l] + matrix[k+2][l+1] + matrix[k+2][l+2]
			if sum >= largest {
				largest = sum
			}
		}
	}
	fmt.Println(largest)
}
