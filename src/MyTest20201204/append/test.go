package main

import "fmt"

type Boy struct {
	Name string
	Age  int
}

func main() {
	var b []Boy

	var b1 = Boy{"aa", 1}
	var b2 = Boy{"bb", 3}

	b = append(b, b2)
	b = append(b, b1)

	fmt.Println(b)

}
