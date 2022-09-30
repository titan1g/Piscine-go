package main

import (
	"fmt"
)

func main() {
	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(Join(toConcat, ":"))
}

func Join(strs []string, sep string) string {
	concat := ""
	for i, res := range strs {
		concat += res
		if i != len(strs)-1 {
			concat += sep
		}
	}
	return concat
}
