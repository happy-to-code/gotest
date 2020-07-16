package main

import "fmt"

func squares(c chan int) {
	for i := 0; i <= 3; i++ {
		num := <-c
		fmt.Println(num * num)
	}
}

// 读缓冲区的操作是渴望式读取. 意味着一旦读操作开始它将读取缓冲区所有数据，直到缓冲区为空。
func main() {
	println("main() started")

	c := make(chan int, 3)
	go squares(c)

	c <- 1
	c <- 2
	c <- 3
	// 因为通道容量是3 如果没有下面的一行代码，通道不会超出容量， main函数也就不会阻塞
	// c <- 4

	fmt.Println("main() stopped")

}
