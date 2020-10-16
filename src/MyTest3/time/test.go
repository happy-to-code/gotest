package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"strings"
	"time"
)

func main() {

	t1 := time.Now().Unix()
	fmt.Println(uint64(t1))
	datetime := time.Now().Format("2006/01/02 15:04:05")
	fmt.Println(datetime)
	fmt.Println(len("d5a7aef9-e7e4-5c9b-bd07-b10cd65958a1"))

	for i := 0; i < 11; i++ {

		ul := uuid.NewV5(uuid.NewV4(), "hello11111111111")
		fmt.Println(len(ul.String()))
		fmt.Println(strings.Replace(ul.String(), "-", "", -1))
	}
}
