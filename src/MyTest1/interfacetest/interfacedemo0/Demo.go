package main

import "fmt"

// 定义一个接口  有两个为实现的方法
type I interface {
	Get() int
	Set(int)
}

// 定义一个结构体  有一个age属性
type SS struct {
	Age int
}

// 结构体SS实现了Get方法
func (s SS) Get() int {
	return s.Age
}

// 结构体SS实现了Set方法
func (s SS) Set(age int) {
	s.Age = age
}

// 函数 参数是接口类型
func f(i I) {
	i.Set(10)
	fmt.Println(i.Get())
}

func main() {
	ss := SS{}
	f(&ss) // pointer
	f(ss)  // value
}
