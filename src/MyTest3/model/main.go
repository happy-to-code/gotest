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

	fmt.Println("==========================")
	var peoples []People
	p1 := People{"aa", 14}
	p2 := People{"bb", 18}
	peoples = append(peoples, p1)
	peoples = append(peoples, p2)
	fmt.Println(peoples)
	fmt.Println(peoples[0].Name)

}
