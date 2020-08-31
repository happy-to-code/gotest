package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s = "10005"
	fmt.Printf("%s\n", strconv.Quote(s))

	b := []byte(s)

	ss := string(b)
	fmt.Printf("%s\n", strconv.Quote(ss))
}
