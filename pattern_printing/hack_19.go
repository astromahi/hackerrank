package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	hash := strings.Repeat("#", n)

	for i := 0; i < n; i++ {
		fmt.Printf("%*s\n", n, hash)
		hash = hash[0 : n-i-1]
	}
}
