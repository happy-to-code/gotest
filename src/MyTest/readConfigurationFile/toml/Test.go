package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	Host     string `toml:"host"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Port     string `toml:"port"`
	QrCode   QrCode
	Body     Body
	Text     []Text
}
type QrCode struct {
	Size    int
	Address string
}
type Body struct {
	LsftSideFontSize   int
	LsftSideFontColor  []int
	bodyTextStartPoint []int
}
type Text struct {
	Key   string
	Value string
}

func main() {
	var conf Config
	if _, err := toml.DecodeFile("E:\\20.06.16Project\\GoTest\\src\\MyTest\\readConfigurationFile\\toml\\test.toml", &conf); err != nil {
		// handle error
	}
	log.Println(conf.Host)
	log.Println("======================")

	fmt.Println("conf.Text:::", conf.Text)
	text := conf.Text
	for one := range text {
		fmt.Println(one)
	}

	fmt.Println("======>11", conf)
	fmt.Println("======>", conf.Body)
	// fmt.Println("======>", conf.QrCode)
	// fmt.Println("3======>", reflect.TypeOf(conf.QrCode))

}

func main1() {
	viper.AddConfigPath("./src/MyTest/readConfigurationFile/toml")
	viper.SetConfigName("test") // 设置配置文件的名字
	viper.SetConfigType("toml") // 设置配置文件类型，可选
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("config file error: %s\n", err)
		os.Exit(1)
	}
	Host := viper.Get("Host")
	fmt.Println("Host:", Host)

}
