// https://www.hackerrank.com/challenges/angry-professor

package main

import (
	"fmt"
)

func main() {
	var t, n, k int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		fmt.Scanf("%d %d", &n, &k)
		array := make([]int, n)

		for j := 0; j < n; j++ {
			fmt.Scan(&array[j])
		}

		peopleLate := 0
		for m := 0; m < n; m++ {
			if array[m] > 0 {
				peopleLate += 1
			}
		}

		if (n - peopleLate) >= k {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}
