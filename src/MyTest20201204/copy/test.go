package main

import "fmt"

type Boy struct {
	Name string
	Age  int
}

func main() {
	var b = Boy{
		Name: "XiaoMing",
		Age:  8,
	}
	fmt.Printf("b1:%+v\n", b)
	var b2 = b
	fmt.Printf("b1:%+v\n", b)
	b.Age = 9
	fmt.Printf("b2:%+v\n", b2)
	fmt.Printf("b:%+v\n", b)

}
