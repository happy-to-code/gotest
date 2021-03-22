package main

import "fmt"

func main() {

	var list []string

	for _, l := range list {
		fmt.Println("---------")
		fmt.Println("l:", l)
	}

	fmt.Println("end:")

}
