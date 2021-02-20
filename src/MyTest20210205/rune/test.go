package main

import "fmt"

func main() {
	str := "hello北京"
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("char= %c\n", r[i])
	}
}
