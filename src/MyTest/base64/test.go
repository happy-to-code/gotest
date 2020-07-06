package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main1() {
	input := []byte("c84ed4594f1b03d2ab04d1fd17a393cc8daee39464191afd998e94ce3d20a3b6")

	// 演示base64编码
	encodeString := base64.StdEncoding.EncodeToString(input)
	fmt.Println("encodeString:", encodeString)

	// 对上面的编码结果进行base64解码
	decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("string(decodeBytes):", string(decodeBytes))

	// fmt.Println()

	// // 如果要用在url中，需要使用URLEncoding
	// uEnc := base64.URLEncoding.EncodeToString([]byte(input))
	// fmt.Println(uEnc)
	//
	// uDec, err := base64.URLEncoding.DecodeString(uEnc)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(string(uDec))
}
