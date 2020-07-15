package main

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("yaml") // 或viper.SetConfigType（“YAML”）
	// 任何需要将此配置放入程序的方法
	var yamlExample = []byte(`
Hacker: true
name: steve
hobbies:
- skateboarding
- snowboarding
- go
clothing:
  jacket: leather
  trousers: denim
age: 35
eyes : brown
beard: true
`)

	viper.ReadConfig(bytes.NewBuffer(yamlExample))

	name := viper.Get("name")
	fmt.Println("name:", name)

	hobbies := viper.Get("hobbies")
	fmt.Println("hobbies:", hobbies)
	fmt.Printf("hobbies type: %T\n", hobbies)

	clothing := viper.Get("clothing")
	fmt.Println("clothing:", clothing)
	fmt.Println(viper.Get("clothing.jacket"))
}
