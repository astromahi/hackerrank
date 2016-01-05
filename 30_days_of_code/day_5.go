package main

import (
	"fmt"
	"math"
)

type input struct {
	a, b, N int
}

func main() {

	var T, seq int
	var in input
	var v []input

	fmt.Scanf("%d\n", &T)
	if T < 0 || T > 500 {
		fmt.Println("testcase constraints failed.")
		return
	}

	for i := 0; i < T; i++ {

		fmt.Scanf("%d%d%d\n", &in.a, &in.b, &in.N)

		if in.a < 0 || in.a > 50 {
			fmt.Println("a constrains failed.")
			return
		}

		if in.b < 0 || in.b > 50 {
			fmt.Println("b constraints failed.")
			return
		}

		if in.N < 1 || in.N > 15 {
			fmt.Println("N constraints failed.")
			return
		}

		v = append(v, in)

	}

	for _, val := range v {
		for j := 0; j < val.N; j++ {
			if j == 0 {
				seq = val.a + (int(math.Pow(2, float64(j))) * val.b)
			} else {
				seq = seq + (int(math.Pow(2, float64(j))) * val.b)
			}

			fmt.Printf("%d ", seq)
		}
		fmt.Print("\n")
	}

}
