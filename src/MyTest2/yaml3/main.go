package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type MyFile struct {
	Dicts []Dicts `yaml:"dicts"`
}
type Dicts struct {
	Item  string `yaml:"item"`
	Label Label  `yaml:"label"`
}
type Label struct {
	Name     string `yaml: "name"`
	Code     int32  `yaml: "code"`
	Describe string `yaml: "describe"`
}

func main() {

	var c MyFile
	// 读取yaml配置文件, 将yaml配置文件，转换struct类型
	conf := c.getConf1()
	fmt.Println(conf)
}

func (myFile *MyFile) getConf1() *MyFile {
	// 应该是 绝对地址
	yamlFile, err := ioutil.ReadFile("E:\\20.06.16Project\\GoTest\\src\\MyTest2\\yaml3\\22.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}

	// err = yaml.Unmarshal(yamlFile, myFile)
	err = yaml.UnmarshalStrict(yamlFile, myFile)

	if err != nil {
		fmt.Println(err.Error())
	}

	return myFile
}
