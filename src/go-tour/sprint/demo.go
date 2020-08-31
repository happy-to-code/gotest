package main

import "fmt"

func main() {
	s := fmt.Sprintf("`"+"json:"+"\"%s\""+"`", "ClientName")
	fmt.Println(s)
}
