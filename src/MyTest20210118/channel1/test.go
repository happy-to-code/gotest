package main

import (
	"strconv"
)

func main() {
	// 创建一个通道
	var result = make(chan string)

	for i := 0; i < 5; i++ {
		go say("Hello World: "+strconv.Itoa(i), result)
	}

	var j = 0
	for s := range result {
		println(s)
		// 关闭通道
		if j >= 4 {
			close(result)
		}
		j++
	}
}

func say(s string, c chan string) {
	c <- s
}
