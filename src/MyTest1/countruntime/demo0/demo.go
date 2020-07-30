package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()

	fmt.Println("Hello")

	for i := 0; i < 1000; i++ {
		fmt.Println(i)
	}
	elapsed := time.Since(t)

	time.Sleep(2 * time.Second)
	elapsed1 := time.Since(t)

	fmt.Println("app elapsed:", elapsed)
	fmt.Println("app elapsed1:", elapsed1)
}
