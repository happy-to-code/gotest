package main

import "fmt"

func main() {
	p := person{name: "张三"}
	p.modify() // 值接收者，修改无效
	fmt.Println(p.String())
}

type person struct {
	name string
}

func (p person) String() string {
	return "the person name is " + p.name
}

func (p person) modify() {
	p.name = "李四"
}
