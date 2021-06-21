package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "http://b2904236d6.zicp.vip:8098/my-oss-test/PID345678_7.json"
	index1 := strings.Index(s, "my-oss-test/")
	fmt.Println(index1)
	index2 := strings.Index(s, ".json")
	content := s[index1+12 : index2]
	fmt.Println(content)
	replace := strings.Replace(content, "_", "/", -1)
	fmt.Println(replace)

	// split := strings.Split(s, "/")
	// fmt.Println(split)
	// fmt.Println(strings.Split(s, "/")[0])
	// fmt.Println(time.Now().Unix())
	// fmt.Println(strings.Contains("product_market_subject_ref", "."))
}
