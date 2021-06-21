package main

import (
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	// 总共有500个协程需要等待执行
	wg.Add(500)

	for i := 0; i < 500; i++ {
		go say("Hello World", i)
	}
	wg.Wait()
}

func say(s string, i int) {
	println(s, "---", i)
	wg.Done()
}
