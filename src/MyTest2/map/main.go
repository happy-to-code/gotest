package main

import "fmt"

type BB struct {
	Name   string
	Amount uint64
}

func main() {
	var bList []BB

	bList = append(bList, BB{"aa", 10})
	bList = append(bList, BB{"bb", 20})
	bList = append(bList, BB{"bb", 5})
	bList = append(bList, BB{"aa", 5})
	bList = append(bList, BB{"cc", 5})
	fmt.Println(bList)

	var m = make(map[string]uint64)
	i := len(m)
	fmt.Println("11::::", i)
	m["ppp"] = 123123
	fmt.Println(m)
	fmt.Println("------------------")
	fmt.Println(m["ppp"])

	for _, b := range bList {
		var isDump = false
		for s, u := range m {
			if b.Name == s {
				m[b.Name] = b.Amount + u
				isDump = true
			}
		}
		if !isDump {
			m[b.Name] = b.Amount
		}
	}

	fmt.Printf("===================")
	fmt.Printf("%v\n", m)

}
