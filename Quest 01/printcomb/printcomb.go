package main

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i <= '9'; i++ {
		for r := '0'; r <= '9'; r++ {
			for s := '0'; s <= '9'; s++ {
				if i < r {
					if r < s {
						z01.PrintRune(i)
						z01.PrintRune(r)
						z01.PrintRune(s)
						if i == '7' && r == '8' && s == '9' {
							z01.PrintRune(10)
						} else {
							z01.PrintRune(44)
							z01.PrintRune(32)
						}
					}
				}

			}
		}
	}
}
