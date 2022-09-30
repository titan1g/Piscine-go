package main

import (
	"fmt"
)

func main() {
	a := 13
	b := 2
	UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
func UltimateDivMod(a *int, b *int) {
	y := *a
	*a = *a / *b
	*b = y % *b
}
