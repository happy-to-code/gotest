package main

import "fmt"

// 带缓冲区channe和不带缓冲区channel
// 带缓冲区channel：定义声明时候制定了缓冲区大小(长度)，可以保存多个数据。
//
// 不带缓冲区channel：只能存一个数据，并且只有当该数据被取出时候才能存下一个数据。
func test(c chan int) {
	for i := 0; i < 100; i++ {
		fmt.Println("send ", i)
		c <- i
	}
}
func main() {
	ch := make(chan int, 10)
	go test(ch)
	for j := 0; j < 100; j++ {
		fmt.Println("get ", <-ch)
	}
}
