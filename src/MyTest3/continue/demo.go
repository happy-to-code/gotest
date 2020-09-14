package main

import "fmt"

func main() {

	var list []string
	list = append(list, "aa")
	list = append(list, "bb")
	list = append(list, "cc")

	fmt.Println("list:", list)
	fmt.Printf("==============\n")
	for _, s := range list {
		if s == "bb" {
			continue
		}
		fmt.Println(s)
	}

}
