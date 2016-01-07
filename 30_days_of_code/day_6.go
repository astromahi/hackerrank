// Printing staircase
package main

import "fmt"

func main() {

	var N int

	_, err := fmt.Scanf("%d", &N)
	if err != nil {
		fmt.Println("something wrong")
		return
	}

	if N < 1 || N > 100 {
		fmt.Println("N constrains failed.")
		return
	}

	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			if (i + j) > N {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}
