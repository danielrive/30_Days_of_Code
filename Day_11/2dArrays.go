/// ##### WARNING #######
/// ##### WARNING #######

/// ## The following code was provided for Hacker Rank,
// some lines have the commen By Daniel, those lines are mine :)

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(readLine(reader), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(6) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)

	}
	/// ##### WARNING #######
   //// THE FOLLOWING CODE IS BY DANIEL R
	var hourglassArr []int32    array to store the sum for each hourglass
	sum := int32(0)
	for j := 0; j <= 3; j++ {
		for p := 0; p <= 3; p++ {
			sum = arr[j][p] + arr[j][p+1] + arr[j][p+2] + arr[j+1][p+1] + arr[j+2][p] + arr[j+2][p+1] + arr[j+2][p+2]
			hourglassArr = append(hourglassArr, sum)
		}

	}
	fmt.Println(hourglassArr)
	result := int32(-1000)
	for k := 0; k < len(hourglassArr); k++ {
		if result < hourglassArr[k] {
			result = hourglassArr[k]
		}

	}
	fmt.Println(result)
	// END DANIEL CODE
	/// ##### WARNING #######
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
