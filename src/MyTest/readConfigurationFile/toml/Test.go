package main

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	Host     string `toml:"host"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Port     string `toml:"port"`
}

func main() {
	var conf Config
	if _, err := toml.DecodeFile("E:\\20.06.16Project\\GoTest\\src\\MyTest\\readConfigurationFile\\toml\\test.toml", &conf); err != nil {
		// handle error
	}
	log.Println(conf.Host)
}
