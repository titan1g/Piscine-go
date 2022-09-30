package main

import (
	"fmt"
)

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}

var brocceta []string

var test int = 0

func Split(s, sep string) []string {
	for test = 0; test <= len(s); test++ {
		if test+len(sep) > len(s) {
			break
		}

	}
	return brocceta
}
