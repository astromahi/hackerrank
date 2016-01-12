package main

import (
	"fmt"
)

var stack []int

func toBinary(n int) {
	if n == 1 {
		stack = append(stack, n)
	} else {
		rem := n % 2
		stack = append(stack, rem)
		toBinary(n / 2)
	}
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
