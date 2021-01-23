package main

// A simple script to print  Hello world and the input(stdin)
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)
	inputsread, _ := scanner.ReadString('\n')
	inputsread = strings.Replace(inputsread, "\n", "", -1)
	fmt.Println("Hello, World.")
	fmt.Println(inputsread)

}
