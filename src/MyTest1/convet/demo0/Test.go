package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	// var s = "fWDoYleTsrtX/FusbTj4EVUdGhUxfilpqaqHqsT2XTY="
	var s = "hello,world我是存证交易2020.08.5"

	// 将字符串转为字节数组
	bytes := []byte(s)
	fmt.Printf("字节数组：%v\n", bytes)
	fmt.Printf("字节数组16进制：%x\n", bytes)
	// 将字符串转为base64
	base64Str := base64.StdEncoding.EncodeToString(bytes)
	fmt.Printf("base64Str：%s\n", base64Str)
	fmt.Printf("base64Str16进制形式: %x\n", base64Str)

	fmt.Println("---------------------------------------")

	// 将字符串 解码成16进制字节数组
	hexStr := hex.EncodeToString([]byte(s))
	fmt.Printf("hexBytes：%s\n", hexStr)

	hexBytes, _ := hex.DecodeString(hexStr)
	fmt.Printf("hexBytes：%s\n", hexBytes)
	fmt.Println("string(hexBytes):::", string(hexBytes))
	fmt.Println(fmt.Sprintf("%s", hexBytes))

}
