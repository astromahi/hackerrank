// Rotating array left - using reverse algorithm
// Complexity O(n)
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

	reverse(array, 0, k-1)
	reverse(array, k, n-1)
	reverse(array, 0, n-1)

	fmt.Println(array)
}

func reverse(arr []int, start, end int) {
	for start < end {
		temp := arr[start]
		arr[start] = arr[end]
		arr[end] = temp
		start++
		end--
		//arr[start], arr[end] = arr[end], arr[start]
	}
}
