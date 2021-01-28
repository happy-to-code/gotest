package main

import (
	"fmt"
	"time"
)

func main() {
	// 字符拼接
	fmt.Println(fmt.Errorf("%v,%s%v%s%v%s", "err", "[", 2, ":", 3, "]"))
	var a  = 5
	t2 := time.Duration(a) * time.Second
	fmt.Println(t2)
	fmt.Println(t2.Seconds())
}
