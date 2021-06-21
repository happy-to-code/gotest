package main

import (
	"fmt"
	"strings"
)

func main() {
	var key = "keyKKKKL"
	var v = "oid123"

	err := fmt.Errorf("%v,%s", "wo shi err", "["+key+":"+v+"]")
	fmt.Printf("%+v\n", err)

	var s = "  ssss  "
	fmt.Println(s, "--")
	fmt.Println(strings.TrimSpace(s), "--")
	fmt.Println(strings.Trim(s, " "), "--")
}
