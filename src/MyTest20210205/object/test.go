package main

import "fmt"

func main() {
	dog := Dog{"旺财", 1, true}
	fmt.Println(dog)
	fmt.Println("----------------")
	dog.changeName()
	fmt.Println(dog)
}

type Dog struct {
	Name   string
	Age    int
	Gender bool
}

func (d *Dog) changeName() {
	d.Name = d.Name + "---"
	fmt.Println("222222222222222:", d)
}
