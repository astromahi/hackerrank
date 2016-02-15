// https://www.hackerrank.com/challenges/matching-anything-but-new-line
// Matching Anything But a Newline

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	var input string

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		input = scanner.Text()
	}

	re := regexp.MustCompile(`...\....\....\....`)
	match := re.FindAllString(input, -1)

	if match != nil {
		fmt.Printf("Number of matches : %d\n", len(match))
	}
}
