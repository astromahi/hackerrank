// https://www.hackerrank.com/contests/30-days-of-code/challenges/day-16-closest-numbers
// Day 16: Sorting!

package main

import (
	"fmt"
	"math"
	"sort"
)

type Lst struct {
	a, b, diff int
}

func main() {
	var n, e int
	var tList []Lst
	var list, result []int

	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &e)
		list = append(list, e)
	}

	// Sorting the list
	sort.Ints(list)

	// Finding minimum abs difference
	mabs := 9223372036854775807
	for j := 0; j < n-1; j++ {
		d := math.Abs(float64(list[j]) - float64(list[j+1]))
		abs := int(d)
		if abs <= mabs {
			mabs = abs
			tList = append(tList, Lst{list[j], list[j+1], mabs})
		}
	}

	// Finding the list elements that generates minimum abs
	for k := 0; k < len(tList); k++ {
		if mabs == tList[k].diff {
			result = append(result, tList[k].a, tList[k].b)
		}
	}

	// Printing the result
	for l := range result {
		fmt.Print(result[l], " ")
	}
}
