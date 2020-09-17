package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var a map[string]json.RawMessage
	a["aa"] = []byte("sss")
	a["bb"] = []byte("bb")
	a["cc"] = []byte("fff")

	fmt.Printf("a:%+v\n", a)
}
