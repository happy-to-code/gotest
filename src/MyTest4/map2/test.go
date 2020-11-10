package main

import (
	"fmt"
	"sync"
)

var m sync.Map

func main() {
	fmt.Println([]byte("dfsdfs"))
	m.Store("a", []byte("dfsdfshahcdnd1232frDF"))
	m.Store("b", "bbb")
	m.Store("c", "ccc")
	fmt.Printf("m %+v\n", m)

	load, ok := m.Load("a")
	fmt.Printf("type %T\n", load.([]uint8))
	fmt.Println(load)
	fmt.Printf("%T\n", load)
	fmt.Println("===>", load.([]uint8))
	fmt.Println(string(load.([]uint8)))
	if ok {
		fmt.Printf("load: %s\n", load)
	}

	load1, ok := m.Load("as")
	if ok {
		fmt.Printf("load1: %s\n", load1)
	}
}
