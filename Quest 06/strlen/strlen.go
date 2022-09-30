package main

import (
	"fmt"
)

func main() {
	l := StrLen("Hello World!")
	fmt.Println(l)
}

func StrLen(ch string) int {
	return len([]rune(ch))
}
