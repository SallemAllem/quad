package piscine

import "github.com/01-edu/z01"

func QuadC(x, y int) {
	if x > 0 && y > 0 {
		for i := 0; i < x; i++ {
			if i == 0 {
				z01.PrintRune('A')
			} else if i == x-1 {
				z01.PrintRune('A')
			} else {
				z01.PrintRune('B')
			}
		}
		if y > 1 {
			for a := 0; a < y-2; a++ {
				z01.PrintRune('\n')
				for b := 0; b < x; b++ {
					if b == 0 || b == x-1 {
						z01.PrintRune('B')
					} else {
						z01.PrintRune(' ')
					}
				}
			}
			z01.PrintRune('\n')
			for i := 0; i < x; i++ {
				if i == 0 {
					z01.PrintRune('C')
				} else if i == x-1 {
					z01.PrintRune('C')
				} else {
					z01.PrintRune('B')
				}
			}
		}
		if x > 0 && y > 0 {
			z01.PrintRune('\n')
		}
	}
}
