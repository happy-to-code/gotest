package main

import (
	"fmt"
	"sync"
)

// goroutine之间的通讯
// goroutine本质上是协程，可以理解为不受内核调度，而受go调度器管理的线程。
// goroutine之间可以通过channel进行通信或者说是数据共享，当然你也可以使用全局变量来进行数据共享。
//
// 示例：使用channel模拟消费者和生产者模式
func Productor(mychan chan int, data int, wait *sync.WaitGroup) {
	mychan <- data
	fmt.Println("product data：", data)
	wait.Done()
}
func Consumer(mychan chan int, wait *sync.WaitGroup) {
	a := <-mychan
	fmt.Println("consumer data：", a)
	wait.Done()
}
func main() {

	datachan := make(chan int, 100) // 通讯数据管道
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		go Productor(datachan, i, &wg) // 生产数据
		wg.Add(1)
	}
	for j := 0; j < 10; j++ {
		go Consumer(datachan, &wg) // 消费数据
		wg.Add(1)
	}
	wg.Wait()
}
