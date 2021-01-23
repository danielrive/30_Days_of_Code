package main

// a code to calculate the factorial using recursion
import (
	"fmt"
)

func main() {
	n := int32(5)
	fmt.Println(factorial(n))
}
func factorial(n int32) int32 {
	if n <= 0 {
		return 1
	}
	return n * factorial(n-1)

}
