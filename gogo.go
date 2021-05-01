package main

import (
	"fmt"
)

func QuadA(x, y int) {
	res := make([][]string, y)

	for i := 0; i < y; i++ {
		if i == 0 || i == y-1 {
			res[i] = handleFirstLastRow(x)
		} else {
			res[i] = handleOtherRows(x)
		}
	}

	for i := 0; i < len(res); i++ {
		row := res[i]
		for j := 0; j < len(row); j++ {
			fmt.Print(row[j])
		}
		fmt.Println()
	}
}

func handleFirstLastRow(x int) []string {
	res := make([]string, x)
	for i := 0; i < x; i++ {
		if i == 0 || i == x-1 {
			res[i] = "o"
		} else {
			res[i] = "-"
		}
	}
	return res
}

func handleOtherRows(x int) []string {
	res := make([]string, x)

	for i := 0; i < x; i++ {
		if i == 0 || i == x-1 {
			res[i] = "|"
		} else {
			res[i] = " "
		}
	}
	return res
}

func main() {
	QuadA(1, 1)
}
