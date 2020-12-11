package main

import "fmt"

func main() {
	var a interface{}
	a = "11"
	b, ok := a.(string)
	fmt.Println(b, "<---")
	if ok {
		fmt.Println(b)
	}
}
