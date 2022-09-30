package main

import (
	"fmt"
)

func main() {
	fmt.Println(MakeRange(5, 10))
	fmt.Println(MakeRange(10, 5))
}
func MakeRange(min, max int) []int {
	if min >= max {
		var answer []int = nil
		return answer
	} else {
		sexe := make([]int, max-min)
		j := 0
		for i := min; i < max; i++ {
			sexe[j] = i
			j++
		}
		return sexe
	}
}
