package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func main() {
	// 声明一个int类型的变量
	num := 3
	fmt.Println("num:", num)

	// 声明一个string类型的变量
	str := "HelloWorld"
	fmt.Println("str:", str)

	// 	现在我们声明一个函数类型的变量
	myAdd := Add
	fmt.Println("myAdd:", myAdd(1, 2))
}
