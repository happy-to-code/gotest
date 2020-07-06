package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func main() {
	viper.AddConfigPath("./src/MyTest/readConfigurationFile/json")
	viper.SetConfigName("config") // 设置配置文件的名字
	viper.SetConfigType("json")   // 设置配置文件类型，可选
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("config file error: %s\n", err)
		os.Exit(1)
	}
	urlValue := viper.Get("mysql.url")
	fmt.Println("mysql url:", urlValue)
	// fmt.Printf("mysql url: %s\n mysql username: %s\n mysql password: %s", viper.Get("mysql.url"), viper.Get("mysql.username"), viper.Get("mysql.password"))
}
