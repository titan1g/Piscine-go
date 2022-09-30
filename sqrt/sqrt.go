package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(3))
}

func Sqrt(nb int) int {
	if nb < 0 || math.MaxInt32 < nb {
		return 0
	}
	for i := 0; i < 101; i++ {
		if nb == i*i {
			nb = i
			break
		} else if i*i > nb {
			nb = 0
			break
		}
	}
	return nb
}
