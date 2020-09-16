package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Boy struct {
	Name string
	Age  int
}

func main() {
	s := strconv.Itoa(1)
	fmt.Printf("%s,%T\n", s, s)
	fmt.Printf("%s,%T", string(s), 1)
	fmt.Println("============")

	var b = Boy{Name: "xiaoa", Age: 11}
	fmt.Println("b:", b)
	prodBytes, _ := json.Marshal(b)
	fmt.Println("prodBytes:", string(prodBytes))
	args1 := []interface{}{prodBytes}
	data1, _ := json.Marshal(args1)
	fmt.Println("data1:", data1)
	fmt.Println("data1:", string(data1))

	fmt.Printf("======================")

	str := "   hello world!   "
	fmt.Println()
	fmt.Println(str, "|")
	str = strings.TrimSpace(str)
	fmt.Println(str, "|")

	fmt.Println(time.Now().Format("3:04:05.000 PM Mon Jan"))            // 2:27:05.702 PM Thu Jul
	fmt.Println(time.Now().Format("2006-01-_2 3:04:05.000 PM Mon Jan")) // 2016-07-14 2:54:11.442 PM Thu Jul
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))               // 2016-07-14 14:54:11.442239513 +0800 CST

	fmt.Println("=============================================")
	var a []byte
	s2 := string(a)
	fmt.Println(s2, "<----")

}
