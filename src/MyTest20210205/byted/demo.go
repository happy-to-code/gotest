package main

import "fmt"

func main() {
	s := "123456789"
	b := []byte(s)
	fmt.Println(len(b))
	fmt.Println(string(b[2:]))
}
