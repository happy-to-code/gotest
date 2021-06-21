package main

import (
	"encoding/json"
	"fmt"
)

const LEDGERID = "ledgerId"

func main() {
	// var m = make(map[string]map[string]int64)

	var m1 = make(map[string]int64)
	m1["aa"] = 1
	m1["bb"] = 2
	m1["cc"] = 3

	fmt.Printf("m1===>%+v\n", m1)
	m1Bytes, err := json.Marshal(m1)
	if err != nil {
		panic(err)
	}
	fmt.Println("m1bytes:", string(m1Bytes))
	var m2 map[string]int64
	json.Unmarshal(m1Bytes, &m2)
	fmt.Printf("%T:::::%+v\n", m2, m2)

	fmt.Println(LEDGERID)
	fmt.Println(LEDGERID + "-------")
}
