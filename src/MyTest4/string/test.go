package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	s := "admins       cadmin secadmin aduadmin"
	fmt.Printf("%q\n", strings.Fields(s))

	fmt.Println("--------------------------")
	fmt.Println(time.Now())
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixNano())
	fmt.Println(time.Now().UnixNano() / 1000000)
	fmt.Println(time.Millisecond)
	fmt.Println("---------------")
	for i := 0; i < 100; i++ {
		// fmt.Println(time.Now().UnixNano() / 1000000)
		fmt.Println(time.Now().UnixNano())
	}

	s1 := "saf"
	s2 := "sae"
	fmt.Println(s1 > s2)

}
