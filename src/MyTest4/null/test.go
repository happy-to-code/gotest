package main

import (
	"fmt"
	"strings"
)

func main() {
	var b []byte

	fmt.Println(b)
	fmt.Println(b == nil)

}

func isNil(s string) bool {
	if strings.TrimSpace(s) == "" {
		return true
	}
	return false
}
