package main

import (
	"fmt"
	"time"
)

func main() {
	i := fun1()
	fmt.Println("i==>", i)
	time.Sleep(100)
}

func fun1() int {
	go fun2()
	fmt.Println("func1:", 1)
	return 1
}
func fun2() {
	fmt.Println("func2:", 2)
}
