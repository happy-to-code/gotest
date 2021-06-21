package main

import (
	"fmt"
	"strings"
)

type People struct {
	Name string
	Age  int
}

func main() {
	var s = "sss.aaa"
	fmt.Println(strings.Contains(s, "."))
	splitStr := strings.Split(s, ".")
	for _, v := range splitStr {
		fmt.Println("v:", v)
	}

	m := make(map[string]interface{})
	m["aa"] = "123"
	m["bb"] = "456"

	m["cc"] = []People{{Name: "xiaoMing", Age: 12}, {Name: "Hua", Age: 18}}

	fmt.Printf("m:%+v\n", m)
	changeValue(m, "cc")
	fmt.Printf("m22:%+v\n", m)

}

func changeValue(m map[string]interface{}, key string) {
	v, has := m[key]
	if has {
		ps, ok := v.([]People)
		if ok {
			for _, p := range ps {
				newName := p.Name + "/99"
				fmt.Println(newName)
				p.Name = newName
			}
		}
	}
}
