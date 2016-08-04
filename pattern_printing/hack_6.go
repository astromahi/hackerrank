package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		for j := n - i; j > 0; j-- {
			fmt.Print("#")
		}
		fmt.Print("\n")
	}
}
