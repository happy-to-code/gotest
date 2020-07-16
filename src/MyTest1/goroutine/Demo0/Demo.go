package main

import (
	"fmt"
)

func main() {
	strChan := make(chan string, 10)
	fmt.Println("strChan:", strChan)
	fmt.Printf("信道可缓冲 %d 个数据\n", cap(strChan))
	fmt.Printf("信道中当前有 %d 个数据\n", len(strChan))
	strChan <- "1"
	fmt.Printf("信道中当前有 %d 个数据\n", len(strChan))
}
