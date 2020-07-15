package main

import (
	"fmt"
	"math"
)

func main() {
	var s = "HelloWorld123!"

	fmt.Println("ssss:", s)
	bytes := []byte(s)

	fmt.Println(fmt.Sprintf("%s", bytes))

	// fmt.Println("bytes:", bytes)
	// fmt.Println("string(bytes):", string(bytes))
	// fmt.Printf("bytesStr %s\n:", bytes)
	// fmt.Printf("bytesV %v\n:", bytes)
	// fmt.Printf("bytes+V %+v\n:", bytes)

	var num uint64 = 1000123456
	fmt.Println("mmm:", num)
	f := float64(num)
	fmt.Println("f:", f)
	f2 := math.Trunc(f/0.000001) * 0.000001
	fmt.Println("f2:", f2)
	fmt.Println("f3:", f/1000000)

	fmt.Println(fmt.Sprintf("value is :%s", f/1000000))
	fmt.Println(fmt.Sprintf("value2 is :%v", f/1000000))
}
