package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type Dictionary struct {
	Name     string
	Code     uint16
	Describe string
}

type Config1 struct {
	DictionaryList []Dictionary
}
type Config struct {
	Server map[string]Dictionary
}

func main() {
	var config Config
	load, _ := toml.DecodeFile("E:\\20.06.16Project\\GoTest\\src\\MyTest2\\const\\type2.toml", &config)

	fmt.Print(load)
	fmt.Println(load.Keys())

	fmt.Println("-----------------")

	s := load.Type("BZ")
	fmt.Println(s)
	fmt.Println(load.Type("BZ.name"))

}
