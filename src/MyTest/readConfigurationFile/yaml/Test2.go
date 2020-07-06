package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func main() {
	var c conf
	conf := c.getConf()
	fmt.Println(conf.Host)
}

// profile variables
type conf struct {
	Host   string `yaml:"host"`
	User   string `yaml:"user"`
	Pwd    string `yaml:"pwd"`
	Dbname string `yaml:"dbname"`
}

func (c *conf) getConf() *conf {
	yamlFile, err := ioutil.ReadFile("E:\\20.06.16Project\\GoTest\\src\\MyTest\\readConfigurationFile\\yaml\\conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}
