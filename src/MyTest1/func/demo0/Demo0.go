package main

import "fmt"

func dd(i func(int, int) int) int {
	return i(1, 2)
}

func main() {
	ee := func(x, y int) int {
		return x + y
	}
	fmt.Println(dd(ee))
}
