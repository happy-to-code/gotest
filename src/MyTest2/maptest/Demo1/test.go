package main

import (
	"fmt"
)

func main() {
	var mymap map[string]string

	mymap = map[string]string{"1a": "Very", "2b": "good", "3c": "day"}

	for k, v := range mymap {
		fmt.Println(k, ":\t", v)

		s := mymap[k]
		fmt.Println(s)
	}

}
