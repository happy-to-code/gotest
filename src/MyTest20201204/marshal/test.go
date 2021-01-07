package main

import (
	"encoding/json"
	"fmt"
)

type Boy struct {
	Name  string `json:"name,omitempty"`
	Age   int64  `json:"age,omitempty"`
	Class string `json:"class,omitempty"`
}

func main() {
	var b = Boy{
		Name: "xiaoMing",
		Age:  12,
	}
	fmt.Printf("b:%+v\n", b)
	boyBytes, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("boyBytesStr:%s\n", boyBytes)
	fmt.Println("--------------------")
	var b2 Boy
	err = json.Unmarshal(boyBytes, &b2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("b2:%+v\n", b2)
	fmt.Println("-->", "<--")
	fmt.Println("-->", b2.Class, "<--")
	fmt.Println(b2.Class == "")
}
