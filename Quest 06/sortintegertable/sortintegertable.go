package main

import (
	"fmt"
)

func main() {
	s := []int{5, 4, 3, 2, 1, 0}
	SortIntegerTable(s)
	fmt.Println(s)
}

func SortIntegerTable(table []int) {
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table); j++ {
			if table[i] < table[j] {
				table[i], table[j] = table[j], table[i]
			}
		}
	}
}
