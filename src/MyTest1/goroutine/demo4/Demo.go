package main

import (
	"fmt"
)

func main() {
	fmt.Println("main() started!")
	c := make(chan string)
	go greet(c)

	c <- "World"

	fmt.Println("main() stopped!")
}

func greet(c chan string) {
	fmt.Println("Hello " + <-c + "!")
}
