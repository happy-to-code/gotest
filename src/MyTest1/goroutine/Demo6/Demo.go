package main

import (
	"fmt"
	"runtime"
)
import "time"

func main() {
	fmt.Println("1run in main goroutine")
	go func() {
		fmt.Println("2run in child goroutine")
		go func() {
			fmt.Println("3run in grand child goroutine")
			go func() {
				fmt.Println("4run in grand grand child goroutine")
			}()
		}()
	}()
	time.Sleep(time.Second)
	fmt.Println("5main goroutine will quit")
}
