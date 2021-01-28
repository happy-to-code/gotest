package main

import (
	"fmt"
	"time"
)

func main() {
	printTime()
}

func printTime() {
	ticker := time.NewTicker(2 * time.Second)
	for _ = range ticker.C {
		fmt.Println(time.Now())
	}
}
func init() {
	println("11111111111111111")
}
func init() {
	go printTime()
}
