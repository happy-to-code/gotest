package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Boy struct {
	Name string
	Age  int
}

func main() {
	s := strconv.Itoa(1)
	fmt.Printf("%s,%T\n", s, s)
	fmt.Printf("%s,%T", string(s), 1)
	fmt.Println("============")

	var b = Boy{Name: "xiaoa", Age: 11}
	fmt.Println("b:", b)
	prodBytes, _ := json.Marshal(b)
	fmt.Println("prodBytes:", string(prodBytes))
	args1 := []interface{}{prodBytes}
	data1, _ := json.Marshal(args1)
	fmt.Println("data1:", data1)
	fmt.Println("data1:", string(data1))
}
