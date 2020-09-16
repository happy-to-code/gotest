package main

import "fmt"

func main() {
	var m map[string]interface{} = make(map[string]interface{}, 5)

	fmt.Println("len(m):", len(m))
	m["aa"] = "xiaoming"
	m["bb"] = "xiaoHua"
	m["cc"] = 3
	m["dd"] = true
	fmt.Println("len(m):", len(m))

	fmt.Println("m:::", m)

	s := getValueFromMap("CC1", "CC1", m)
	fmt.Println("s:", s)
	fmt.Println(s == nil)
	// fmt.Println("s111:", s.(int))

	fmt.Println("last____")
}

func getValueFromMap(k1, k2 string, m map[string]interface{}) interface{} {
	var inter interface{}

	inter = m[k1]
	if inter == nil {
		inter = m[k2]
	}
	return inter
}
