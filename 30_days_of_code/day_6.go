// Printing staircase
package main

import "fmt"

func main() {

	var N int

	fmt.Scanf("%d", &N)

	if N < 1 || N > 100 {
		fmt.Println("N constrains failed.")
		return
	}

	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			if (i + j) > N {
				fmt.Printf("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Printf("\n")
	}
}
