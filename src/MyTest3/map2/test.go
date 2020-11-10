package main

import "fmt"

func main() {
	m := make(map[string]interface{}, 5)

	m["aa"] = 11
	m["bb"] = 22
	m["cc"] = "fg"
	fmt.Printf("m1:%+v\n", m)
	m["aa"] = -123
	m["dd"] = -123
	fmt.Printf("m1:%+v\n", m)
	fmt.Printf("m1:%+v\n", m["cc"].(string))
	fmt.Println(m["cc1"] == nil)
	// fmt.Printf("m1:%+v\n", m["cc1"].(string))

	fmt.Println("=========================")

}
