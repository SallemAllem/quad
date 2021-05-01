package main

import "github.com/01-edu/z01"

func check(x int) {
	for i := 0; i < x; i++ {
		if i == 0 || i == x-1 {
			z01.PrintRune('o')
		} else {
			z01.PrintRune('-')
		}
	}
}

func QuadA(x, y int) {
	check(x)
	if y > 1 {
		for a := 0; a < y-2; a++ {
			z01.PrintRune('\n')
			for b := 0; b < x; b++ {
				if b == 0 || b == x-1 {
					z01.PrintRune('|')
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
		check(x)
	}
	z01.PrintRune('\n')
}
func main() {
	QuadA(5, 3)
	QuadA(10, 30)
}
