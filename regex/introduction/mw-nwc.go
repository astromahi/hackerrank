// https://www.hackerrank.com/challenges/matching-word-non-word
// Matching Word & Non-Word Character

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

	re := regexp.MustCompile(`\w\w\w\W\w\w\w\w\w\w\w\w\w\w\W\w\w\w`)
	matched := re.MatchString(input)
	fmt.Println(matched)
}
