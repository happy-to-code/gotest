package main

import (
	"MyTest20201204/redis/conn"
	"fmt"
)

func main() {
	// isSet := SetEx("k", "abc", 20)
	// fmt.Println("isSet:", isSet)
	value := conn.Get("k")
	fmt.Printf("value:%s\n", value)

	value1 := conn.Get("kk")
	fmt.Printf("value1:%s\n", value1)

}
