package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const BufferSize = 100

func main() {
	file, err := os.Open("./bb.json")
	// file, err := os.Open("E:\\20.06.16Project\\GoTest\\src\\MyTest3\\readjson\\bb.json")
	// file, err := os.Open("E:\\20.06.16Project\\GoTest\\src\\MyTest3\\readjson\\12.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("content:", string(content))
}
