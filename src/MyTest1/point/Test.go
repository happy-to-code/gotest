package main

import "fmt"

func main() {
	var s = "HelloWorld"
	fmt.Println("s值：", s)
	point := &s
	fmt.Println("point:", point)

	fmt.Println("*point:", *point)
	s = *point + "123"
	fmt.Println("s修改后值：", s)
	fmt.Println("*point修改后值:", *point)

}
