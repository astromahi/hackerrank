package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	var hash string
	for i := 0; i < n; i++ {
		hash += "#"
		fmt.Println(hash)
	}
}
