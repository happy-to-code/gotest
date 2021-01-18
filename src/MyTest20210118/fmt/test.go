package main

import "fmt"

func main() {
	// 字符拼接
	fmt.Println(fmt.Errorf("%v,%s%v%s%v%s", "err", "[", 2, ":", 3, "]"))
}
