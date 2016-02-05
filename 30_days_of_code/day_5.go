// https://www.hackerrank.com/contests/30-days-of-code/challenges/day-5-loops
// Day 5: Loops!

package main

import (
	"fmt"
	"math"
)

func main() {

	var t, a, b, n, seq int

	fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		fmt.Scanf("%d%d%d", &a, &b, &n)
		seq = a + b
		fmt.Print(seq, " ")
		for j := 1; j < n; j++ {
			seq = seq + int(math.Pow(2, float64(j)))*b
			fmt.Printf("%d ", seq)
		}
		fmt.Print("\n")
	}

}
