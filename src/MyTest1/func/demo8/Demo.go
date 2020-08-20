package main

import "fmt"

type CalculateFun func(a, b int) int

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Calculate(a, b int, calculate CalculateFun) int {
	return calculate(a, b)
}

func main() {
	a, b := 5, 3

	// 	相加
	add := Calculate(a, b, Add)
	fmt.Println("add:", add)
	// 	相减
	sub := Calculate(a, b, Sub)
	fmt.Println("sub:", sub)
}
