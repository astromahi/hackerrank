package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		space := n - i - 1
		for space >= 0 {
			fmt.Print(" ")
			space--
		}

		hash := i
		for hash >= 0 {
			fmt.Print("#")
			hash--
		}

		fmt.Print("\n")
	}
}
