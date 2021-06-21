package main

import "fmt"

func main() {
	defer name(1)
	defer name(2)
	fmt.Println("11111111111111111111111")
}

func name(i int) {
	fmt.Println("name========>", i)
}
