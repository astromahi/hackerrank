// Rotating array left - using extra memory
// Complexity O(n*k)
// Space complexity O(k)
package main

import "fmt"

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)

	array := make([]int, n)
	tempArray := make([]int, k)
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}

	k = k % n
	if k == 0 || k == n {
		fmt.Println(array)
		return
	}

	for j := 0; j < k; j++ {
		tempArray[j] = array[j]
	}

	fmt.Println("-----------Steps----------")
	for m := 0; m < k; m++ {
		for p := 0; p < n-1; p++ {
			array[p] = array[p+1]
			fmt.Println(array)
		}
	}
	fmt.Println("-----------Steps----------")

	for q := 0; q < k; q++ {
		array[n-k+q] = tempArray[q]
	}
	fmt.Println(array)
}
