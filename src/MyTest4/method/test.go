package main

import "fmt"

func main() {
	var boy = B{
		Name: "",
		Age:  10,
	}
	name := boy.getName()
	fmt.Println(name)
}
