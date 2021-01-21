package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var testcases int
	var arrInfo []string
	info := make(map[string]string)
	fmt.Scan(&testcases)

	// reader := bufio.NewReader(os.Stdin)

	// for i := 0; i < testcases; i++ {
	// 	text, _ := reader.ReadString('\n')
	// 	arrInfo = strings.Split(text, " ")
	// 	info[arrInfo[0]] = arrInfo[1]
	// 	arrInfo = nil
	// }
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input2 := scanner.Text()
		if testcases != 0 {
			arrInfo = strings.Split(input2, " ")
			info[arrInfo[0]] = arrInfo[1]
			arrInfo = nil
			testcases--
		} else {
			if input2 != "" {
				_, isThere := info[input2]
				if isThere {
					fmt.Printf("%s=%s\n", input2, info[input2])
				} else {
					fmt.Println("Not found")
				}
			} else {
				break
			}
		}
	}

}
