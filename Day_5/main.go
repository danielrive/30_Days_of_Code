package main

import (
	"fmt"
)

func main() {
	n := 5

	for i := 1; i < 11; i++ {
		fmt.Printf("%d x %d = %d\n", n, i, i*n)
	}

}
