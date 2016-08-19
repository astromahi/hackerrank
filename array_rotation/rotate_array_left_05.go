// Rotating array left - using array index manipulation
// Complexity O(n)
// Space complexity O(n)
package main

import "fmt"

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)

	array := make([]int, n)
	tempArray := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}

	k = k % n
	if k == 0 || k == n {
		fmt.Println(array)
		return
	}

	for j := 0; j < n; j++ {
		//temp1 := array[j]
		index := computeIndex(j+k, n)
		//temp2 := array[index]
		tempArray[j] = array[index]
		//array[index] = temp1
	}

	fmt.Println(tempArray)
}

func computeIndex(k, n int) int {

	for i := 1; k < 0; i++ {
		k = k + (i * n)

	}

	return k % n
}
