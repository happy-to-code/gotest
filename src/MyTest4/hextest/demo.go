package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	// 将base64形式  转换成16进制
	s := "8IvFbQKLxn4HaIHzWdehvzADjjgxMj7jJWBdaji3WNA="

	decodeString, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		fmt.Printf("%v", err)
	}
	toString := hex.EncodeToString(decodeString)
	fmt.Printf("%s\n", decodeString)
	fmt.Println("====>", toString)
}
