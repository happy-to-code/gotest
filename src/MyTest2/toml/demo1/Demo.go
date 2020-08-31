package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type Dict struct {
	Name     string
	Code     int
	Describe string
}
type Dicts struct {
	DictList []Dict
}

func main() {
	var dicts Dicts

	if _, err := toml.DecodeFile("D:\\200707Go\\gotest\\src\\MyTest2\\toml\\demo1\\135.toml", &dicts); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("dicts:%+v\n", dicts)

	for _, dict := range dicts.DictList {
		fmt.Println(dict.Name)
		fmt.Println(dict.Code)
		fmt.Println(dict.Describe)
	}

}
