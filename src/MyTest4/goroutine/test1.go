package main

import (
	"fmt"
	"time"
)

func addele(a []int, i int) {
	defer func() { // 匿名函数捕获错误
		err := recover()
		if err != nil {
			fmt.Println("add ele fail:", i)
		}
	}()
	a[i] = i
	fmt.Println(a)
}

func main() {
	arry := make([]int, 4)
	for i := 0; i < 10; i++ {
		go addele(arry, i)
	}
	time.Sleep(time.Second * 2)
}
