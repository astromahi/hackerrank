package main

import (
	"fmt"
	"math"
)

type testCase struct {
	a, b, N int
}

func main() {

	var T, seq int
	var tc testCase
	var testCases []testCase

	fmt.Scanf("%d\n", &T)
	if T < 0 || T > 500 {
		fmt.Println("testcase constraints failed.")
		return
	}

	for i := 0; i < T; i++ {

		fmt.Scanf("%d%d%d\n", &tc.a, &tc.b, &tc.N)

		if tc.a < 0 || tc.a > 50 {
			fmt.Println("a constrains failed.")
			return
		}

		if tc.b < 0 || tc.b > 50 {
			fmt.Println("b constraints failed.")
			return
		}

		if tc.N < 1 || tc.N > 15 {
			fmt.Println("N constraints failed.")
			return
		}

		testCases = append(testCases, tc)

	}

	for _, c := range testCases {
		for j := 0; j < c.N; j++ {
			if j == 0 {
				seq = c.a + (int(math.Pow(2, float64(j))) * c.b)
			} else {
				seq = seq + (int(math.Pow(2, float64(j))) * c.b)
			}

			fmt.Printf("%d ", seq)
		}
		fmt.Print("\n")
	}

}
