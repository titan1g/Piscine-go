package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	result := Map(IsPrime, a)
	fmt.Println(result)
}

func Map(f func(int) bool, a []int) []bool {
	counter := 0
	for range a {
		counter++
	}
	res_arr := make([]bool, counter)
	for i, val := range a {
		res_arr[i] = f(val)
	}
	return res_arr
}

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 2; i <= nb/2; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
