package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	hash := "#"
	for i := 0; i < n; i++ {
		fmt.Printf("%*s\n", n, hash)
		hash += "#"
	}
}
