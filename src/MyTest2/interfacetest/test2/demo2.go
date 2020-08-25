package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	fmt.Println("汪汪汪……")
	return ""
}

type Cat struct {
}

func (c *Cat) Speak() string {
	fmt.Println("喵喵喵……")
	return ""
}
func main() {

	animalList := []Animal{Dog{}, &Cat{}}

	for _, animal := range animalList {
		animal.Speak()
	}
	fmt.Println("------------------")
	var dog Dog
	dog.Speak()
	fmt.Println("============")
	dog1 := new(Dog)
	dog1.Speak()

}
