package main

import (
	"fmt"
	"strings"
)

func main() {
	m1 := make(map[string]interface{}, 0)
	m1["aa"] = "AA"
	m1["bb"] = "BB"
	m1["cc"] = "CC"

	// 参数
	m2 := make(map[string]interface{}, 0)
	m2["aa"] = "AAA---"
	m2["bb"] = false
	m2["dd"] = "DDD--"
	fmt.Println(m1)

	fmt.Println("--------------------------")
	for k, v := range m2 {
		var isExist = false
		for key, _ := range m1 {
			if strings.EqualFold(k, key) {
				m1[key] = v
				isExist = true
			}
		}
		if !isExist {
			m1[k] = v
		}
	}
	fmt.Println(m1)
}
