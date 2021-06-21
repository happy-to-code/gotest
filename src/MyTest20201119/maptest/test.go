package main

import (
	"fmt"
	"strings"
)

func main() {
	var m = make(map[string]interface{})
	m["aa"] = "1.txt,2.pdf"
	m["bb"] = "BBBB"
	m["cc"] = []string{"11.txt", "22.pdf"}
	fmt.Println("m:", m)
	fmt.Println("================")
	FieldStringToList(m)
	fmt.Println("m:", m)

}

// 将字符串 更新成list
func FieldStringToList(m map[string]interface{}) {
	var keys = []string{"aa",
		"bb",
		"cc",
		"account_thaw_certificate",
		"account_associated_acct_certificates",
		"account_qualification_certification_file"}
	for _, k := range keys {
		value := m[k]
		if value != nil {
			switch value.(type) {
			case string:
				splitStrSlice := strings.Split(value.(string), ",")
				m[k] = splitStrSlice
			case []string:
				continue
			}
		}
	}
}
