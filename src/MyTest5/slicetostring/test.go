package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = []string{"AA", "bbb", "cc"}
	fmt.Println(s)

	trimStr := strings.Trim(fmt.Sprint(s), "[]")
	fmt.Println(trimStr)

	replace := strings.Replace(strings.Trim(fmt.Sprint(s), "[]"), " ", ",", -1)
	fmt.Println(replace)
}
