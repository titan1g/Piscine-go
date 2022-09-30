package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}

func BasicAtoi(s string) int {
	i, err := strconv.Atoi(s)

	if err == nil {
		return i
	} else {
		return 0
	}
}
