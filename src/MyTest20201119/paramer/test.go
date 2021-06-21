package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a = []int{11, 3, 44, 5}

	typeName := reflect.TypeOf("a").Name()
	fmt.Println("Type:", typeName)
	fmt.Println("string" == typeName)

	fmt.Println("a:", a)
	var b = []int{88}
	b = append(b, a...)
	fmt.Println("b:", b)
	fmt.Println("----------------------------------")
	nums := []interface{}{11, 3, 5, 6}
	TestArgs(99, nums...)
	fmt.Printf("main 中 addr of slice %p\n", nums)
}

func TestArgs(first int, arg ...interface{}) {
	fmt.Println(first, "\t", arg)
	fmt.Println(reflect.TypeOf(arg))
	fmt.Printf("TestArgs 中 addr of slice %p\n", arg)
}
