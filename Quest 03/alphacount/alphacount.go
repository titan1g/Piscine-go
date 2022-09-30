package main

import (
	"fmt"
)

func main() {
	s := "Hello 78 World!  44555 /"
	nb := AlphaCount(s)
	fmt.Println(nb)
}

func AlphaCount(str string) int {
	count := 0
	for _, s := range str {
		if isAlpha(s) {
			count++
		}

	}
	return count
}

func isAlpha(alpha rune) bool {
	for a := 'a'; a <= 'z'; a++ {
		if alpha == a {
			return true
		}
	}
	for a := 'A'; a <= 'Z'; a++ {
		if alpha == a {
			return true
		}
	}
	return false
}
