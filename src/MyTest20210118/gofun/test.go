package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		ticker := time.NewTicker(2 * time.Second)
		for range ticker.C {
			fmt.Println(time.Now())
		}
	}()

	time.Sleep(time.Second*10)
}
