package main

import (
	"fmt"
)

func test() {

	for ch := 97; ch <= 122; ch++ {
		fmt.Printf("%c   ", ch)
	}
	fmt.Println()
}
