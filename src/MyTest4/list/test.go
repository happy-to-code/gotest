package main

import "fmt"

func main() {
	var olds = []string{"0000000000000000", "a", "b", "c", "d"}
	fmt.Println(olds)
	var s = []string{"0000000000000000", "a1", "b1", "a"}
	fmt.Println(s)

	var count int = 0

	for _, s2 := range s {
		var flag = true
		for _, old := range olds {
			if s2 == old {
				flag = false
				continue
			}
		}
		if flag {
			count++
		}
	}

	fmt.Println(count)
}
