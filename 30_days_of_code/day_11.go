//https://www.hackerrank.com/contests/30-days-of-code/challenges/day-11-more-review
//Day 11: 2D-Arrays + More Review!
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
	for k := 1; k < 5; k++ {
		for l := 1; l < 5; l++ {
			sum = matrix[k-1][l-1] + matrix[k-1][l] + matrix[k-1][l+1] + matrix[k][l] + matrix[k+1][l-1] + matrix[k+1][l] + matrix[k+1][l+1]
			if sum >= largest {
				largest = sum
			}
		}
	}
	fmt.Println(largest)
}
