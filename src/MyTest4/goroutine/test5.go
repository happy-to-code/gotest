package main

import (
	"fmt"
)

// 循环管道
//
// 需要注意的是：
//
// 使用range循环管道，如果管道未关闭会引发deadlock错误。
// 如果采用for死循环已经关闭的管道，当管道没有数据时候，读取的数据会是管道的默认值，并且循环不会退出。
func main() {
	// 创建一个带缓冲区的通道
	mychannel := make(chan int, 10)
	for i := 0; i < 10; i++ {
		mychannel <- i
	}
	close(mychannel) // 关闭管道--->不影响后面通道的使用
	fmt.Println("data lenght: ", len(mychannel))
	for v := range mychannel { // 循环管道
		fmt.Println(v)
	}
	fmt.Printf("data lenght:  %d", len(mychannel))
}
