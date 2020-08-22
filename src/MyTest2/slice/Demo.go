package main

import "fmt"

func main() {
	var strList []string
	strList = append(strList, "aa")
	strList = append(strList, "bb")
	strList = append(strList, "cc")

	fmt.Println(strList)
	fmt.Println("11" == "11")

	var req struct {
		M              uint8
		SubAccountType string
		Format         string
		Args           []string
	}

	fmt.Println(req)

}
