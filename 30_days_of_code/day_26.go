// https://www.hackerrank.com/contests/30-days-of-code/challenges/day-26-testing-part-1-plus-implementations
// Day 26: Testing Part I + Implementations!

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	var fine int
	var at, et string

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		at = scanner.Text()
	}

	if scanner.Scan() {
		et = scanner.Text()
	}

	dateLayout := `2 1 2006`
	actualDate, _ := time.Parse(dateLayout, at)
	expectedDate, _ := time.Parse(dateLayout, et)

	ay, am, ad := actualDate.Date()
	ey, em, ed := expectedDate.Date()

	if ay == ey {
		if am == em {
			if ad <= ed {
				fine = 0
			} else {
				fine = 15 * (ad - ed)
			}
		} else if am > em {
			fine = 500 * (int)(am-em)
		} else {
			fine = 0
		}
	} else if ay > ey {
		fine = 10000
	} else {
		fine = 0
	}

	fmt.Println(fine)

}
