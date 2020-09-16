package main

import (
	"errors"
	"fmt"
)

func main() {
	notFound := errors.New("404 not found")
	fmt.Println(notFound)

	// 也可以使用fmt包中包装的Errorf来添加
	fmt.Println(fmt.Errorf("404: page %v is not found", "index.html"))
}
