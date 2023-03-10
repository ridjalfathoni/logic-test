package main

import (
	"fmt"
	"strings"
)

func soal1() {
	x := []int{1, 4, 1, 2, 4, 2, 7, 9, 4, 5}
	yAxis := ""
	xAxis := "  "

	for i := 9; i >= 1; i-- {
		yAxis = fmt.Sprintf("%d ", i)

		for j := 0; j < len(x); j++ {
			if x[j] >= i {
				yAxis += "# "
			} else {
				yAxis += "- "
			}
			if i == 9 {
				xAxis += fmt.Sprintf("%d ", j)
			}
		}
		fmt.Println(yAxis)
	}

	fmt.Println("  " + strings.Repeat("- ", len(x)))
	fmt.Print(xAxis)
}

func main() {
	soal1()
}
