package main

import (
	"fmt"
)

func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(ConcatParams(test))
}

func ConcatParams(bite []string) string {
	str := ""
	for i, res := range bite {
		str += string(res)
		if i != len(bite)-1 {
			str += "\n"
		}
	}
	return str
}
