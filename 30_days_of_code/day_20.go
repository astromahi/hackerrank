// https://www.hackerrank.com/contests/30-days-of-code/challenges/day-20-review-plus-more-string-methods
// Day 20: Review + More String Methods!

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {

	var line string
	var words []string
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		line = scanner.Text()
	}

	line = strings.Trim(line, " ![,?.\\_'@+]")

	//regex := regexp.MustCompile(`[^a-zA-Z\s!\[,?.\\_'@+\]]`)
	//line = regex.ReplaceAllLiteralString(line, " ")
	//fmt.Println(line)
	regex := regexp.MustCompile(`[\s!\[,?.\\_'@+\]]+`)
	data := regex.Split(line, -1)

	for _, val := range data {
		w := strings.TrimSpace(val)
		if w != "" {
			words = append(words, w)
		}
	}

	fmt.Println(len(words))
	for i := range words {
		fmt.Println(words[i])
	}
}
