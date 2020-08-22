package main

import "fmt"

func main() {
	var boy = Boy{
		Name: "xiaoMing",
		Age:  17,
	}
	fmt.Printf("boy:%+v\n", boy)
	fmt.Println(boy.Name == "xiao")
	fmt.Println(boy.Name == "xiaoMing")
	fmt.Println(boy.Name == "")

}

type Boy struct {
	Name string
	Age  uint64
}
