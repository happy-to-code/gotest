package main

import "fmt"

func main() {
	var m = make(map[string]interface{})
	m["aa"] = "AA"
	m["bb"] = "BB"

	// fmt.Printf("m:%+v\n", m)
	// valueFromMap, b := getValueFromMap("cc", "aa", m)
	// fmt.Println(valueFromMap, ":", b)

	fmt.Printf("m:%+v\n", m)
	var m2 map[string]interface{}
	m2 = m
	fmt.Printf("m2:%v\n", m2)
	fmt.Println("=======================")
	m["aa"] = "aAa"
	fmt.Printf("m:%+v\n", m)
	fmt.Printf("m2:%v\n", m2)
}

func getValueFromMap(k1, k2 string, m map[string]interface{}) (interface{}, bool) {
	inter, has := m[k1]
	if !has {
		inter, has = m[k2]
	}
	return inter, has

}
