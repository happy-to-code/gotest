package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a = make([]int, 3)
	a[0] = 2
	a[1] = 3
	a[2] = 4
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))

	var b = make([]int, 2)
	b[0] = 20
	b[1] = 21

	ints := append(b, a...)
	fmt.Println(ints)
	// 品牌商的方式
}
