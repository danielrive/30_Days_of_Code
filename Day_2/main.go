package main

// A simple script to sum int, float and concatenate string

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	i2, _ := strconv.ParseUint(scanner.Text(), 10, 64)
	scanner.Scan()
	d2, _ := strconv.ParseFloat(scanner.Text(), 64)
	scanner.Scan()
	s2 := scanner.Text()
	fmt.Println(i + i2)
	fmt.Printf("%.1f\n", d+d2)
	fmt.Println(s + s2)

}
