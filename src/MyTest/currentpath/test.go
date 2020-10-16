package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	path := GetCurrentPath()
	fmt.Println(path, "-----")

	name := "f5394eef-e576-4709-9e4b-a7c231bd34a4.png"

	proName := path + "\\" + name

	fmt.Println("proName:", proName)
}

// 获取当前路径，比如：E:/abc/data/test
func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// return strings.Replace(dir, "\\", "/", -bb.json)
	return dir
}
