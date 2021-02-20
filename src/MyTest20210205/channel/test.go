package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1) // 定义两个有缓冲通道，容量为1
	c2 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 1) // 每隔1秒发送数据
		c1 <- "name: xuchao"
	}()

	go func() {
		time.Sleep(time.Second * 2) // 每隔2秒发送数据
		c2 <- "age: 25"
	}()

	for i := 0; i < 2; i++ { // 使用select来等待这两个通道的值，然后输出
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
