package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you?"))
}

func SplitWhiteSpaces(str string) []string {
	tab := strings.Fields(str)
	return tab
}
