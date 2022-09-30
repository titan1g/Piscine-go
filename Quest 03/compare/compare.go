package main

import (
	"fmt"
)

func main() {
	fmt.Println(Compare("Hello!", "Hello!"))
	fmt.Println(Compare("Salut", "lut!"))
	fmt.Println(Compare("01a", "01"))
}
func Compare(str, str1 string) int {
	len_str := len(str)
	len_str1 := len(str1)
	min := min_size(len_str, len_str1)

	for i := 0; i < min; i++ {
		if str[i] < str1[i] {
			return -1
		}
		if str1[i] < str[i] {
			return 1
		}
	}
	if len_str < len_str1 {
		return -1
	}
	if len_str1 < len_str {
		return 1
	}
	return 0
}

func min_size(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
