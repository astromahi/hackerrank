// https://www.hackerrank.com/contests/30-days-of-code/challenges/day-10-binary-numbers\
package main

import (
	"fmt"
)

var stack []int

func toBinary(n int) {
	for n > 1 {
		rem := n % 2
		stack = append(stack, rem)
		n = n / 2
	}
	stack = append(stack, rem)
}

func main() {
	var t, n int
	fmt.Scanf("%d\n", &t)

	for i := 0; i < t; i++ {
		fmt.Scanf("%d\n", &n)
		toBinary(n)
		for i := 1; i <= len(stack); i++ {
			fmt.Print(stack[len(stack)-i])
		}
		fmt.Print("\n")
		stack = nil
	}
}
