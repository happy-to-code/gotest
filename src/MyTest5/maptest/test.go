package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	m := make(map[string]interface{})
	m["AA"] = "aa"
	m["BB"] = "bb"
	m["CC"] = "cc"

	fmt.Println("m==>", m)
	fmt.Println(m["AAA"] == nil)
	currentPath := GetCurrentPath()
	fmt.Println(currentPath)
}

// 获取当前项目路径
func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
