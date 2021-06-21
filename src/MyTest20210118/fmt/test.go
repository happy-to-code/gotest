package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(time.Now().Unix())
		time.Sleep(time.Second)
		fmt.Println(time.Now().Unix())
	}
}
