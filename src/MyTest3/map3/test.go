package main

import (
	"fmt"
	"strings"
)

func main() {
	m := make(map[string]interface{}, 5)
	m["基本信息"] = map[string]interface{}{
		"年龄": 28,
		"姓名": "刘德华",
	}

	v1, b := getValueFromMap(m, "基本信息.年龄")
	fmt.Println(":::", v1.(int), "  ", b)

	v2, b := getValueFromMap(m, "基本信息.年龄1")
	fmt.Println(":::", v2, "  ", b)
	fmt.Println(":::", v2 == nil)

}

func getValueFromMap(m map[string]interface{}, path string) (interface{}, bool) {
	segs := strings.Split(path, ".")
	for i, seg := range segs {
		if m != nil && len(m) > 0 {
			if v, ok := m[seg]; ok {
				if i == len(segs)-1 {
					return m[seg], true
				} else {
					if m, ok = v.(map[string]interface{}); !ok {
						return nil, false
					}
				}
			}
		}
	}
	return nil, false
}
