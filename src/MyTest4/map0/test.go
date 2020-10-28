package main

import (
	"fmt"
	"sort"
)

func main() {
	var m = map[string]int{
		"unix":         0,
		"python":       1,
		"go":           2,
		"javascript":   3,
		"testing":      4,
		"philosophy":   5,
		"startups":     6,
		"productivity": 7,
		"hn":           8,
		"reddit":       9,
		"C++":          10,
	}
	fmt.Printf("%+v\n", m)
	v, has := m["C++"]
	if !has {
		fmt.Printf("=====%v\n", has)
	} else {
		fmt.Printf("==========----%v\n", v)
	}
	fmt.Println("__________________________________________________")
	fromMap, b := getValueFromMap("reddit0", "reddit", m)
	fmt.Println("AAAAAAAAAAAAAAAA:", fromMap, b)

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", m[k])
	}
}
func getValueFromMap(k1, k2 string, m map[string]int) (interface{}, bool) {
	inter, has := m[k1]
	if !has {
		inter, has = m[k2]
	}
	return inter, has
}
