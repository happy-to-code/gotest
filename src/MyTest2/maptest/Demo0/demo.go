package main

import "fmt"

func main() {
	m := make(map[string]float64, 2)

	fmt.Println("len(m):", len(m))

	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	fmt.Println("len(m):", len(m))

	fmt.Println("m:", m)
	fmt.Println(m["c"])
	fmt.Println(m["c1"])

	fmt.Println("------------------")
	var a float64 = 90
	fmt.Println(a)
	var b uint64 = uint64(a)
	fmt.Println(b)
}
