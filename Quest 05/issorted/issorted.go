package main

import (
	"fmt"
)

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}

	result1 := IsSorted(f, a1)
	result2 := IsSorted(f, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}

func IsSorted(f func(a, b int) int, tab []int) bool {
	res := 1
	res2 := 1
	res3 := 1
	for k, v := range tab {
		if k != len(tab)-1 {
			if f(v, tab[k+1]) < 0 {
				res++
			}
			if f(v, tab[k+1]) > 0 {
				res2++
			}
			if f(v, tab[k+1]) == 0 {
				res3++
			}
		}
	}
	return res == len(tab) || res2 == len(tab) || res3 == len(tab)
}
