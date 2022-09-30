package main

import (
	"fmt"
)

func main() {
	fmt.Println(RecursivePower(4, 3))
}

func RecursivePower(nb, power int) int {
	if power < 0 {
		return 0
	} else if 1 == power {
		return nb
	} else if 0 == power {
		return 1
	} else {
		return nb * RecursivePower(nb, power-1)
	}
}
