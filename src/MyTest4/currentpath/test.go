package main

import (
	"log"
	"os"
)

func main() {
	path := GetCurrentPath()
	println(path)
} // 获取当前项目路径
func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
