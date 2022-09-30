package main

import (
	"github.com/01-edu/z01"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	ForEach(PrintNbr, a)

}

func PrintNum(num int) {
	i := '0'
	if num == 0 {
		z01.PrintRune('0')
		return
	}
	for j := 1; j <= num%10; j++ {
		i++
	}
	for j := -1; j >= num%10; j-- {
		i++
	}
	if num/10 != 0 {
		PrintNum(num / 10)
	}
	z01.PrintRune(i)
	return
}

func ForEach(f func(int), a []int) {

	for _, i := range a {
		f(i)

	}
}

func PrintNbr(n int) {

	if n < 0 {
		z01.PrintRune('-')
	}
	PrintNum(n)
}
