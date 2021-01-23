package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	n := 100
	countCurrent := 0
	countFinal := 0
	nInBinary := strconv.FormatInt(int64(n), 2)
	arrOnes := strings.Split(nInBinary, "0")
	for i := range arrOnes {
		countCurrent = len(arrOnes[i])
		if countCurrent > countFinal {
			countFinal = countCurrent
		}
	}
	fmt.Println(countFinal)
}
