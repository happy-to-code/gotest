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

}
