package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(IsUpper("HELLO"))
	fmt.Println(IsUpper("HELLO!"))

}

func IsUpper(str string) bool {
	for _, res := range str {
		if !unicode.IsUpper(res) {
			return false
		}
	}
	return true
}
