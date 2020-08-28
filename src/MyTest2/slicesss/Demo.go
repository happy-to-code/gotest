package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Name string
	Age  uint32
}

func main() {
	var peoples []People

	p1 := People{
		Name: "asdf",
		Age:  13,
	}
	p2 := People{
		Name: "dgrgf",
		Age:  18,
	}
	peoples = append(append(peoples, p1), p2)
	fmt.Printf("%+v\n", peoples)

	fmt.Println("===============")
	bytes, _ := json.Marshal(peoples)
	fmt.Println(bytes)
	fmt.Println("----------------")
	var ps []People
	json.Unmarshal(bytes, &ps)
	fmt.Printf("ps:%+v\n", ps)

}
