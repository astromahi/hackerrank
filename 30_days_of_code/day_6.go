// https://www.hackerrank.com/contests/30-days-of-code/challenges/day-6-lets-review
// Day 6: Let's Review!

package main

import "fmt"

func main() {

	var n int
	fmt.Scanf("%d", &n)

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if (i + j) > n {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}

}
