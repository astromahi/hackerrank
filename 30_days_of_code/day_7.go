package main

import "fmt"

func main() {
	var n, v int
	var array [1000]int

	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println("something went wrong")
		return
	}

	if n < 1 || n > 1000 {
		fmt.Println("n constraints failed")
		return
	}

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &v)
		array[i] = v
	}

	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", array[n-i])
	}

}
