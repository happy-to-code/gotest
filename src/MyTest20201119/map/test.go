package main

import "fmt"

func main() {
	m := make(map[string]interface{})
	m["aa"] = "123"

	v, _ := m["vv"]
	fmt.Println(v)
	fmt.Println("v:", v.(string))

}
