package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个无缓冲的通道
	pipline := make(chan int)

	go func() {
		num := <-pipline
		fmt.Printf("接收到的数据是: %d", num)
	}()

	go func() {
		fmt.Println("准备发送数据: 100")
		pipline <- 100
	}()

	// 主函数sleep，使得上面两个goroutine有机会执行
	time.Sleep(1)
}
