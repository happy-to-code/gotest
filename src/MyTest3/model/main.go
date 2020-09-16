package main

import "fmt"

type People struct {
	Name   string
	Gender int
}
type Boy struct {
	*People
	Address string
}

func main() {
	var p = &People{
		Name:   "xiaoMing",
		Gender: 1,
	}
	fmt.Println(p)
	fmt.Println(*p)

	var b = Boy{People: p, Address: "suzhou"}

	fmt.Printf("b:%+v\n", b)
	fmt.Printf("b1:%+v\n", b.People)
	fmt.Printf("b2:%+v\n", b.People.Name)

}
