// https://www.hackerrank.com/contests/30-days-of-code/challenges/day-20-review-plus-more-string-methods
// Day 20: Review + More String Methods!

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func splitWord(data []byte, atEOF bool) (int, []byte, error) {

	var accum []byte

	for i, b := range data {
		if b == ' ' || b == '\n' || b == '\t' || b == '\r' || b == '!' || b == '[' || b == ',' || b == '?' || b == '.' || b == '\\' || b == '_' || b == '\'' || b == '@' || b == '+' || b == ']' {

			if len(accum) != 0 {
				return i, accum, nil
			}
		} else {
			accum = append(accum, b)
		}
	}

	if atEOF && len(accum) != 0 {
		return len(data), accum, nil
	}

	return 0, nil, nil
}

func main() {

	var words []string

	scanner := bufio.NewScanner(os.Stdin)
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = splitWord(data, atEOF)
		//fmt.Println("EOF: ", atEOF, "advance: ", advance, "token: ", string(token), "err: ", err)
		return
	}
	scanner.Split(split)

	for scanner.Scan() {
		words = append(words, strings.TrimSpace(scanner.Text()))
	}

	fmt.Println(len(words))
	for i := range words {
		fmt.Println(words[i])
	}

}
