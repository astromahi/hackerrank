// https://www.hackerrank.com/contests/30-days-of-code/challenges/day-25-running-time
// Day 25: Running Time and Complexity!

package main

import (
	"fmt"
	"math"
)

func main() {

	var t, n int

	fmt.Scanf("%d\n", &t)

	for i := 0; i < t; i++ {
		fmt.Scanf("%d\n", &n)

		if n == 1 {
			fmt.Println("Not prime")
			continue
		}

		count := 0
		for j := 1; j <= int(math.Sqrt(float64(n))); j++ {
			if n%j == 0 {
				count++
			}
		}

		if count == 1 {
			fmt.Println("Prime")
		} else {
			fmt.Println("Not prime")
		}

	}
}
