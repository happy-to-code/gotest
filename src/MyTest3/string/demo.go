package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	var s = "al_bal_SnBLwhWVVQHiSJ1CkfNUet1aos1vUVCGHFEADjHL2VQpUotSmJg_3688:1000000000"

	addr := getAddr(s)
	fmt.Println(addr)

	var ss = "1000000000"
	parseInt, _ := strconv.ParseInt(ss, 10, 64)
	fmt.Println(parseInt)

	fmt.Println("=====================")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Duration(1) * time.Second)
		fmt.Println(i, "===============")

	}
}

func getAddr(s string) string {
	index0 := strings.Index(s, "_")

	s2 := s[index0+1:]

	index2 := strings.Index(s2, "_")

	s3 := s2[:index2]
	return s3
}
