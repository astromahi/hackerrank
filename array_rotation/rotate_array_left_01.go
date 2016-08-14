// Rotating array left using rotation cases
// Complexity O(k*n)
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

	var j, temp int
	fmt.Println("-----------Steps----------")
	for m := 0; m < k; m++ {
		temp = array[0]
		for j = 0; j < n-1; j++ {
			array[j] = array[j+1]
			fmt.Println(array)
		}
		array[j] = temp
	}
	fmt.Println("-----------Steps----------")
	fmt.Println(array)
}
