package main

import "fmt"

type Animal interface {
	Speak()
	Eat()
}

// -------------------------
type Dog struct {
	Name string
	Age  int
}

type Cat struct {
	Name string
	Age  int
}

// -------------------------
func (d Dog) Speak() {
	fmt.Println(d.Name, d.Age, "wangwang.....")
}
func (d Dog) Eat() {
	fmt.Println(d.Name, d.Age, "吃骨头。。。")
}

// -------------------------
func (c Cat) Speak() {
	fmt.Println(c.Name, c.Age, "miaomiao.....")
}
func (c Cat) Eat() {
	fmt.Println(c.Name, c.Age, "吃鱼.....。。。")
}

func main() {
	animals := []Animal{Dog{"旺财", 3}, Cat{"小花", 6}}
	for _, a := range animals {
		a.Speak()
		a.Eat()
		fmt.Println("///////////////////////////////")
	}
}
