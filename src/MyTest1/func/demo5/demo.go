package main

import "fmt"

func main() {
	aa := 10
	fmt.Println("=======>", &aa)
	test0(&aa)
	fmt.Println("-----------------------------")
	test1(aa)
}

func test0(a *int) {
	fmt.Println("point A:", a)
	fmt.Println("point A:", *a)

	// *a = 11
}

func test1(a int) {
	fmt.Println("value A:", a)
	fmt.Println("value A:", &a)
}
