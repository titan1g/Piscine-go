package main

import (
	"fmt"
)

func main() {
	fmt.Println(isAlpha("Hello! How are you?"))
	fmt.Println(isAlpha("HelloHowareyou"))
	fmt.Println(isAlpha("What's this 4?"))
	fmt.Println(isAlpha("Whatsthis4"))
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

// 	if alpha == a {
// 		return true
// 	}
// }
// for a := 'A'; a <= 'Z'; a++ {
// 	if alpha == a {
// 		return true
// 	}
// }
// return false
