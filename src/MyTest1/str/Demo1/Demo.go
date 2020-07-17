package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "HelloWorld"
	s1 := s[5:]
	fmt.Println("newS:", s1)

	lastIndexOfd := strings.LastIndex(s, "d")
	fmt.Println("lastIndexOfd:", lastIndexOfd)

	s2 := s[:5]
	fmt.Println("newS:", s2)

}
