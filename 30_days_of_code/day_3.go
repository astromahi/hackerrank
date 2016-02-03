// https://www.hackerrank.com/contests/30-days-of-code/challenges/day-3-if-statements
// Day 3: If-Else Statements!

package main

import (
	"fmt"
)

func main() {

	var n int
	fmt.Scanf("%d", &n)

	if n < 0 || n > 100 {
		fmt.Println("constraints failed")
		return
	}

	mod := n % 2

	if mod != 0 {
		fmt.Println("Weird")
	} else if (mod == 0) && (n <= 5) {
		fmt.Println("Not Weird")
	} else if (mod == 0) && (n <= 20) {
		fmt.Println("Weird")
	} else if (mod == 0) && (n > 20) {
		fmt.Println("Not Weird")
	} else {
		fmt.Println("Weird")
	}

}
