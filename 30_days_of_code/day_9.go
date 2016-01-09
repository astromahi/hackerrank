package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func gcd(x, y int) int {
	if x == y {
		return x
	} else {
		return gcd(max(x, y)-min(x, y), min(x, y))
	}
}

func main() {
	var x, y int
	fmt.Scanf("%d\n%d\n", &x, &y)
	fmt.Println(gcd(x, y))
}
