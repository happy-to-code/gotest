package main

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var a = 5
	fmt.Println("a:==================", a)

	var b int
	if a == 5 {
		for i := 0; i < 10; i++ {
			// 休眠500毫秒
			time.Sleep(time.Duration(1) * time.Second)
			b = rand.Intn(10)
			fmt.Println("[b====]:", b)
			if b == a {
				a = b
				break
			}
		}
	}

	fmt.Println("out B:", b)
	fmt.Println("out A:", a)

	fmt.Println("===================")

	toString := hex.EncodeToString(nil)
	fmt.Println("toString:", toString)

}
