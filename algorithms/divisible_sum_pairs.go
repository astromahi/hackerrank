// https://www.hackerrank.com/challenges/divisible-sum-pairs

package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)

	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}

	count := 0
	for j := 0; j < n; j++ {
		for l := j + 1; l < n; l++ {
			if (array[j]+array[l])%k == 0 {
				count++
			}
		}
	}

	fmt.Println(count)
}
