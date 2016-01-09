package main

import (
	"bufio"
	"fmt"
	"os"
	//"regexp"
)

func main() {

	var n int
	var name, pno string
	var contacts map[string]string

	scanner := bufio.NewScanner(os.Stdin)
	contacts = make(map[string]string)

	fmt.Scanf("%d\n", &n)

	for i := 0; i < n; i++ {

		if scanner.Scan() {
			name = scanner.Text()
		}

		//matched, _ := regexp.MatchString(`^[a-z]+(\s[a-z]+)?`, name)
		//if !matched {
		//fmt.Println("name constraints failed: ", matched)
		//return
		//}

		if scanner.Scan() {
			pno = scanner.Text()
		}

		//matched, _ = regexp.MatchString(`^[1-9][\d]{7}$`, pno)
		//if !matched {
		//fmt.Println("pno constraints failed.")
		//return
		//}

		contacts[name] = pno
	}

	for scanner.Scan() {
		name = scanner.Text()
		no, ok := contacts[name]
		if ok {
			fmt.Printf("%s=%s\n", name, no)
		} else {
			fmt.Println("Not found")
		}
	}

}
