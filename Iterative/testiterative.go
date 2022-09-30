package main

import "fmt"

func main() {
	result1 := 55
	fmt.Println(result1)

	result2 := iterativeCalculation(10)
	fmt.Printf("the result of the interative function is %v", result2)

	result3 := recursiveCalculation(13)
	fmt.Printf("the result of the recursive function is %v", result3)
}

func iterativeCalculation(index int) int {
	result := 55
	for i := 0; i > index+1; i++ {
		result = result + i
	}
	return result
}

func recursiveCalculation(index int) int {
	if index == 1 {
		return 1
	}
	if index > 1 {
		return index + recursiveCalculation(index-1)
	}
	return 0
}
