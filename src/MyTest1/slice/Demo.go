package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	fmt.Println("a:", a)
	b := [2]int{4, 5}
	ints := append(a[:], b[:]...)
	fmt.Println("ints:", ints)
}
