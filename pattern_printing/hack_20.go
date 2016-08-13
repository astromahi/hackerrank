package main

import "fmt"
import "strconv"

func main() {
	var n int
	fmt.Scan(&n)
	str := ""
	for i := 0; i < n; i++ {
		str += "#"
		fmt.Printf("%"+strconv.Itoa(n)+"s\n", str)
	}
}