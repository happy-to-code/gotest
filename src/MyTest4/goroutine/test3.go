package main

import (
	"fmt"
	"time"
)

// 示例二：通过channel实现goroutine之间的同步。
//
// 实现方式：通过channel能在多个groutine之间通讯，当一个goroutine完成时候向channel发送退出信号,等所有goroutine退出时候，
// 利用for循环channe去channel中的信号，若取不到数据会阻塞原理，等待所有goroutine执行完毕，
// 使用该方法有个前提是你已经知道了你启动了多少个goroutine。
func cal3(a int, b int, Exitchan chan bool) {
	c := a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
	time.Sleep(time.Second * 2)
	Exitchan <- true
}

func main() {

	Exitchan := make(chan bool, 10) // 声明并分配管道内存
	for i := 0; i < 10; i++ {
		go cal3(i, i+1, Exitchan)
	}
	for j := 0; j < 10; j++ {
		<-Exitchan // 取信号数据，如果取不到则会阻塞
	}
	close(Exitchan) // 关闭管道
}
