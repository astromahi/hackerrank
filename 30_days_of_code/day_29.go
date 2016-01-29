// https://www.hackerrank.com/contests/30-days-of-code/challenges/day-29-look-at-everything-weve-learned
// Day 29: Look at Everything We've Learned!

package main

import (
	"fmt"
	"math"
)

func abs(c byte) float64 {
	return math.Abs(float64(c))
}

func main() {

	var t int
	var s string
	fmt.Scanf("%d\n", &t)

	for i := 0; i < t; i++ {
		fmt.Scanf("%s", &s)
		count := 0
		for j := 0; j < len(s)-1; j++ {
			//fmt.Printf("%s - %s, %s - %s\n", string(s[j+1]), string(s[j]), string(s[len(s)-2-j]), string(s[len(s)-1-j]))
			//fmt.Println("Abs1: ", math.Abs(float64(s[j+1]-s[j])), "Abs2: ", math.Abs(math.Abs(float64(s[len(s)-2-j])-s[len(s)-1-j])))
			if math.Abs(abs(s[j+1])-abs(s[j])) == math.Abs(abs(s[len(s)-2-j])-abs(s[len(s)-1-j])) {
				count++
			}
		}
		//fmt.Println(count)
		if count == len(s)-1 {
			fmt.Println("Funny")
		} else {
			fmt.Println("Not Funny")
		}

	}
}
