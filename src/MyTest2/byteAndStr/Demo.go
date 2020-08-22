package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	var s = "我是1234！￥"

	fmt.Printf("s::%s\n", s)
	b := []byte(s)
	fmt.Printf("b::%v\n", b)
	str := string(b)
	fmt.Println(str)
	fmt.Printf("------------------")
	fmt.Printf("bbbb:%x\n", b)
	fmt.Println("bbbb1111:", fmt.Sprintf("%x", b))
	fmt.Println("bbbb2222:", hex.EncodeToString(b))

}
