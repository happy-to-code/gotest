package main

import "fmt"

type Boy struct {
	Name string
	Age  int
}

func main() {
	var b1 = Boy{"xiaoMing", 3}
	funcName(b1)

}

func funcName(b1 Boy) {

	var b2 = b1

	fmt.Println("b1:", b1)
	fmt.Println("b2:", b2)
	b1.Age = 4
	fmt.Println("_____________")
	fmt.Println("b2::::", b2)
	fmt.Println("b1:", b1)
}
