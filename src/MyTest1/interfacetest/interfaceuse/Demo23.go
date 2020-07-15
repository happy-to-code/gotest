package main

import "fmt"

type Type struct {
	name string
}

type PType struct {
	name string
}

type Inter interface {
	post()
}

// 接收者非指针
func (t Type) post() {
	fmt.Println("POST")
}

// 接收者是指针
func (t *PType) post() {
	fmt.Println("POST")
}

func main() {
	var it Inter

	// var it *Inter //接口不能定义为指针
	pty := new(PType)
	ty := Type{name: "type"}

	fmt.Printf("pty type:%T,ty Type:%T\n", pty, ty)

	it = ty   // 将变量赋值给接口，OK
	it.post() // 接口调用方法，OK
	it = pty  // 把指针变量赋值给接口，OK
	it.post() // 接口调用方法，OK

}
