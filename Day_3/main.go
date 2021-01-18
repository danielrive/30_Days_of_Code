package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "math"
)

func main() {
	solve(100,10,20)
}

// Complete the solve function below.
func solve(meal_cost float64, tip_percent int32, tax_percent int32) {
    tip_percent2:=float64(tip_percent)/100
    tax_percent2:=float64(tax_percent)/100
    cost := meal_cost+meal_cost*tip_percent2+meal_cost*tax_percent2
    fmt.Println(math.Round((cost)))

}
