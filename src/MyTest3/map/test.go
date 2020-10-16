package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := make(map[string]interface{}, 5)

	m["aa"] = 11
	m["bb"] = true
	m["cc"] = "fg"
	fmt.Printf("m1:%+v\n", m)
	mBytes, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("mBytes:%s\n", mBytes)
	fmt.Println("mBytes2:", mBytes)
	s := "[" + string(mBytes) + "]"
	fmt.Println(s)
	fmt.Println([]byte(s))
	fmt.Println(string([]byte(s)))
	// marshal, _ := json.Marshal(s)
	// fmt.Printf("marshal:%s\n", marshal)

}
