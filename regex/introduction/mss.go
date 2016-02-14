// https://www.hackerrank.com/challenges/matching-specific-string
// Matching Specific String

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

	regex := regexp.MustCompile(`hackerrank`)
	matchedString := regex.FindAllString(input, -1)

	if matchedString != nil {
		fmt.Printf("Number of matches : %d\n", len(matchedString))
	}
}
