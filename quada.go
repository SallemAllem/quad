package main

import (
	"fmt"
	"github.com/01-edu/z01"
	"os"
	"strconv"
)

func QuadA(x, y int) {
	if x > 0 && y > 0 {
		for i := 1; i <= y; i++ {
			for j := 1; j <= x; j++ {
				if (i == 1 && (j == 1 || j == x)) || (i == y && (j == 1 || j == x)) {
					z01.PrintRune('o')
				} else if (j == 1 && (i != 1 || i != y)) || (j == x && (i != 1 || i != y)) {
					z01.PrintRune('|')
				} else if i != 1 && j != 1 && i != y && j != x {
					z01.PrintRune(' ')
				} else {
					z01.PrintRune('-')
				}
			}
			z01.PrintRune('\n')
		}
	}
}
func main() {
	args := os.Args[1:]
	int1, err := strconv.Atoi(args[0])
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	int2, err := strconv.Atoi(args[1])
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	QuadA(int1, int2)
}
