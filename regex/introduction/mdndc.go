// https://www.hackerrank.com/challenges/matching-digits-non-digit-character
// Matching Digits & Non-Digit Characters

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

	re := regexp.MustCompile(`\d\d\D\d\d\D\d\d\d\d`)
	matched := re.MatchString(input)

	fmt.Println(matched)
}
