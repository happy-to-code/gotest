package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a = 3
	fmt.Printf("%T,%v\n", a, a)

	b := strconv.Itoa(a)
	fmt.Printf("%T,%v\n", b, b)

	c := string(a)
	fmt.Printf("%T,%v\n", c, c)

	fmt.Println("=====================")
	fmt.Println(11 == 1)
}
