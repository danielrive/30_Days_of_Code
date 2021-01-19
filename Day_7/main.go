package main

import (
	"fmt"
)

func main() {
	n := 4
	arr := [5]int{5, 3, 1, 2}
	for j := n - 1; j >= 0; j-- {
		fmt.Printf("%d ", arr[j])
	}

}
