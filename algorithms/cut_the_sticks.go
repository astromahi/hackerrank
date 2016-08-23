// https://www.hackerrank.com/challenges/cut-the-sticks

package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	array := make([]int, n)
	smallest := int(math.Pow(2.0, 31.0))

	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
		if array[i] < smallest {
			smallest = array[i]
		}
	}

	for {
		count := 0
		tempSmall := int(math.Pow(2.0, 31.0))
		for j := 0; j < n; j++ {
			if array[j] > 0 {
				array[j] -= smallest
				count++
			}

			if array[j] > 0 && array[j] < tempSmall {
				tempSmall = array[j]
			}
		}

		if count == 0 {
			break
		} else {
			fmt.Println(count)
		}

		smallest = tempSmall
	}
}
