package main

import (
	"fmt"
	"reflect"
)

type Dog struct {
	Name string
	Age  int
}

func main() {
	var dog = Dog{
		Name: "xiao hua",
		Age:  3,
	}
	fmt.Println("type:", reflect.TypeOf(dog))
	fmt.Printf("Dog:%+v\n", dog)

	fmt.Println("value:", reflect.ValueOf(dog))
	// fmt.Println("convertValue:", reflect.ValueOf(dog).Interface().(string))
}
