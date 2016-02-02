// https://www.hackerrank.com/contests/30-days-of-code/challenges/day-2-arithmetic
// Day 2: Arithmetic!

package main

import (
	"fmt"
	"math"
)

func round(f float64) float64 {
	return math.Floor(f + .5)
}

func main() {

	var m float64
	var t, x int

	fmt.Scanf("%f\n%d\n%d", &m, &t, &x)

	tip := (m * float64(t)) / 100
	tax := (m * float64(x)) / 100

	price := m + tip + tax
	fmt.Printf("The final price of the meal is $%.f.", round(price))
}
