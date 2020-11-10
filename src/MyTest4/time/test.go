package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	var s = "d3232883f194c480adb0c3febd561771963b5f355f7bd5b31f7e0cfee434c39f"
	bytes, _ := hex.DecodeString(s)
	decodeString := base64.StdEncoding.EncodeToString(bytes)

	fmt.Printf("%s\n", decodeString)

	// var ss = "0yMog/GUxICtsMP+vVYXcZY7XzVfe9WzH34M/uQ0w58="
	var ss = "d3232883f194c480adb0c3febd561771963b5f355f7bd5b31f7e0cfee434c39f"

	i, err := hex.DecodeString(ss)
	if err != nil {
		fmt.Println(err)
	}
	toString := hex.EncodeToString(i)
	fmt.Println(toString)
}
