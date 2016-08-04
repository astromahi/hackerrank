package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}

		for k := 0; k <= i; k++ {
			fmt.Print("#")
		}

		fmt.Print("\n")
	}

}
