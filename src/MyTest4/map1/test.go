package main

import (
	"bytes"
	"fmt"
)

func main() {
	var lit []interface{}
	lit = append(lit, "11", "22", "33")
	m := make(map[string]interface{}, 0)
	m["aa"] = lit
	m["bb"] = "llsdkflsd"

	fmt.Printf("%v\n", m)
	convertMapValueListToString(m, "aa")
	fmt.Printf("%v\n", m)
}

func convertMapValueListToString(m map[string]interface{}, key string) {
	v, has := m[key]
	if !has {
		return
	}
	list, ok := v.([]interface{})
	if !ok {
		return
	}

	var buf bytes.Buffer
	for _, str := range list {
		fmt.Fprintf(&buf, "%v,", str)
	}
	fmt.Println("--->", buf.String())
	buf.Truncate(buf.Len() - 1)
	m[key] = buf.String()
}
