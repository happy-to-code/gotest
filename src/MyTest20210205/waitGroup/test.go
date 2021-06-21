package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// 设置计数器的数量
	wg.Add(10)

	// 开N个后台打印线程
	for i := 0; i < 10; i++ {

		go func() {
			// 执行完方法后将计数器的数量-1
			defer wg.Done()

			fmt.Println("你好, 世界", "--", i)
		}()
	}
	// 阻塞代码的执行  直达计数器的值减少到0
	wg.Wait()
}
