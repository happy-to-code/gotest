package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var s = "teestoo1/0"
	sList := strings.Split(s, "/")
	fmt.Println(sList)
	fmt.Println(sList[0])
	fmt.Println("==========================")
	fmt.Println(s + "/" + strconv.Itoa(int(0)))
	fmt.Println(sList[0] + "/" + strconv.Itoa(int(0)))
}
