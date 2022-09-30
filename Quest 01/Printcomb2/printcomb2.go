package main

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	for i := 0; i <= 99; i++ {
		for r := i + 1; r <= 99; r++ {
			ki := i / 10
			ui := i % 10
			kr := r / 10
			ur := r % 10
			z01.PrintRune(rune('0' + ki))
			z01.PrintRune(rune('0' + ui))
			z01.PrintRune(32)
			z01.PrintRune(rune('0' + kr))
			z01.PrintRune(rune('0' + ur))
			if i == 98 && r == 99 {
				z01.PrintRune(10)
			} else {
				z01.PrintRune(44)
				z01.PrintRune(32)
			}
		}
	}
}
