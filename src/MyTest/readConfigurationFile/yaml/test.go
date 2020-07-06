package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

func main1() {
	viper.SetConfigName("config_yaml") // 把json文件换成yaml文件，只需要配置文件名 (不带后缀)即可
	viper.AddConfigPath(".")           // 添加配置文件所在的路径
	// viper.SetConfigType("json")       //设置配置文件类型
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("config file error: %s\n", err)
		os.Exit(1)
	}

	viper.WatchConfig() // 监听配置变化
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置发生变更：", e.Name)
	})

	urlValue := viper.Get("mysql.url")
	fmt.Println("mysql url:", urlValue)
	fmt.Printf("mysql url: %s\nmysql username: %s\nmysql password: %s", viper.Get("mysql.url"), viper.Get("mysql.username"), viper.GetString("mysql.password"))
}
