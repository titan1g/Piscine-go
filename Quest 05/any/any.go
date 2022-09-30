package main

import (
	"fmt"
)

func main() {
	a1 := []string{"Hello", "how", "are", "you"}
	a2 := []string{"This", "is", "4", "you"}

	result1 := Any(IsNumeric, a1)
	result2 := Any(IsNumeric, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}

func Any(f func(string) bool, arr []string) bool {
	isNumeric := false
	for _, val := range arr {
		if f(val) {
			isNumeric = true
		}
	}
	return isNumeric
}

func IsNumeric(s string) bool {
	for _, i := range s {
		if i < '0' || i > '9' {
			return false
		}
	}
	return true
}
