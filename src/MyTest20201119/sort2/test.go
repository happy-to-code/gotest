package main

import (
	"fmt"
	"sort"
)

func main() {
	var a = []int{1, 5, 6, 9, 0, 21}
	sort.Ints(a)
	fmt.Println("a:", a)
}
