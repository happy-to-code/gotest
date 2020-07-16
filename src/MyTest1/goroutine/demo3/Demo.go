package main

import "fmt"

func fibonacci(myChan chan int) {
	n := cap(myChan)
	x, y := 1, 1
	for i := 0; i < n; i++ {
		myChan <- x
		x, y = y, x+y
	}
	// 记得 close 信道
	// 不然主函数中遍历完并不会结束，而是会阻塞。
	close(myChan)
}

func main() {
	pipeline := make(chan int, 10)

	go fibonacci(pipeline)

	for k := range pipeline {
		fmt.Println(k)
	}
}
