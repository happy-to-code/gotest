package main

import (
	"encoding/json"
	"fmt"
)

type Boy struct {
	Name string
	Age  int
}

func main() {
	jsonTest()
}

func jsonTest() {
	var b = Boy{"xiaoMing", 12}

	fmt.Println("b:", b)

	bJsonBytes, _ := json.Marshal(b)
	fmt.Println("bJsonBytes:", bJsonBytes)
}
