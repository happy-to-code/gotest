package main

import (
	"fmt"
	"time"
)

func Producer(queue chan<- int) {
	for i := 0; i < 100; i++ {
		queue <- i // 写入
		fmt.Println("create----> :", i)
	}
}

func Consumer(queue <-chan int) {
	for {
		v := <-queue // 读出
		fmt.Println("receive:", v)
	}
}

func main() {
	queue := make(chan int, 10)
	go Producer(queue)
	go Consumer(queue)
	time.Sleep(10 * time.Second)
}
