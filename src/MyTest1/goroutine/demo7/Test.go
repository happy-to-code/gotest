package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}

	c := make(chan int)

	go sum(s[:2], c) // 3
	go sum(s[2:], c) // 12
	x, y := <-c, <-c
	fmt.Printf("x: %d, y: %d,x/y: %d", x, y, x/y)
}

func sum(s []int, c chan int) {
	sum := 0

	for _, v := range s {
		sum += v
	}
	c <- sum
}
