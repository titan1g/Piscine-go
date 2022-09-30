package main

import (
	"fmt"
	"math"
)

func main() {
	arg := 4
	fmt.Println(IterativeFactorial(arg))
}
func IsPrime(nb int) bool {
	if nb < 2 {
		return false
	} else if 2 == nb || 3 == nb {
		return true
	} else {
		limit := math.Sqrt(float64(nb))
		for i := 2; i <= int(limit); i++ {
			if nb%i == 0 {
				return false
			}
		}
		return true
	}
}
func IterativeFactorial(x int) int {
	if 0 > x {
		return 0
	} else if 0 == x || 1 == x {
		return 1
	} else {
		res := 1
		for i := 1; i <= x; i++ {
			res *= i
			if math.MaxInt32 < res {
				res = 0
				break
			}
		}
		return res
	}
}

func FindNextPrime(nb int) int {
	if IsPrime(nb) {
		return nb
	} else {
		n := nb + 1
		for IsPrime(n) == false {
			n++
		}
		return n
	}
}
