// https://www.hackerrank.com/contests/30-days-of-code/challenges/day-20-review-plus-more-string-methods
// Day 20: Review + More String Methods!

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var line string
	var word []rune
	var words []string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line = scanner.Text()

	for _, b := range line {
		if b == ' ' || b == '\n' || b == '\t' || b == '\r' || b == '!' || b == '[' || b == ',' || b == '?' || b == '.' || b == '\\' || b == '_' || b == '\'' || b == '@' || b == '+' || b == ']' {
			if len(word) != 0 {
				words = append(words, string(word))
				word = nil
			}
		} else {
			//fmt.Println(string(b))
			word = append(word, b)
		}
	}

	if len(word) != 0 {
		words = append(words, string(word))
	}

	fmt.Println(len(words))
	for i := range words {
		fmt.Println(words[i])
	}
}
