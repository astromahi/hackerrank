// Rotating array left - Bubble sort technique
// Complexity O(n*k)
// Space complexity O(1)
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

	k = k % n
	if k == 0 || k == n {
		fmt.Println(array)
		return
	}

	fmt.Println("-----------Steps----------")
	for j := 0; j < k; j++ {
		for l := 0; l < n-1; l++ {
			temp := array[l]
			array[l] = array[l+1]
			array[l+1] = temp
			fmt.Println(array)
		}
	}
	fmt.Println("-----------Steps----------")
	fmt.Println(array)
}
