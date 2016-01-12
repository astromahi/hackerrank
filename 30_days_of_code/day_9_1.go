package main

import "fmt"

func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

func main() {
	var x, y int
	fmt.Scanf("%d\n%d\n", &x, &y)
	fmt.Println(gcd(x, y))
}
