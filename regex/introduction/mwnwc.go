// https://www.hackerrank.com/challenges/matching-whitespace-non-whitespace-character
// Matching Whitespace & Non-Whitespace Character

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

	re := regexp.MustCompile(`\S\S\s\S\S\s\S\S`)
	matched := re.MatchString(input)

	fmt.Println(matched)
}
