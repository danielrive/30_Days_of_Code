package main

import (
	"fmt"
)

func main() {
	var testcases int
	var input string
	var inputs []string
	fmt.Scan(&testcases)
	for i := 0; i < testcases; i++ {
		fmt.Scan(&input)
		inputs = append(inputs, input)
	}
	for j := range inputs {
		fmt.Println(divide(inputs[j]))
	}
	//Enter your code here. Read input from STDIN. Print output to STDOUT
}

func divide(s string) string {
	oddpart := ""
	evenpart := ""

	for i := 0; i < len(s); i++ {
		if i == 0 || i%2 == 0 {
			evenpart = evenpart + string(s[i])
		} else {
			oddpart = oddpart + string(s[i])
		}
	}

	return evenpart + " " + oddpart
}
