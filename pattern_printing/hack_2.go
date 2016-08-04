package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	array := make([]string, n)
	for i := range array {
		array[i] = " "
	}

	for i := 0; i < n; i++ {
		array[i] = "#"
		fmt.Println(strings.Join(array, ""))
	}
}
