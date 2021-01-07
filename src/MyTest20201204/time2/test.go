package main

import (
	"fmt"
	"time"
)

func main() {

	timeUnix := time.Now().Unix()
	timeUnixNano := time.Now().UnixNano()

	fmt.Println(timeUnix)
	fmt.Println(timeUnixNano)

}
