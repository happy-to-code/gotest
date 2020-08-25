package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	configFile := flag.String("c", "", "Configuration File")
	flag.Parse()
	if *configFile == "" {
		fmt.Println("\n\nUse -h to get more information on command line options\n")
		fmt.Println("You must specify a configuration file")
		os.Exit(1)
	}

	err := Initialize(*configFile)
	if err != nil {
		fmt.Printf("Error reading configuration: %s\n", err.Error())
		os.Exit(1)
	}
	// 打印配置值

	fmt.Println(MustGetString("server.maps"))

	mode := MustGetString("server.mode")

	baseURL := MustGetString(mode + ".mode")

}
func Initialize(fileName string) error {
	splits := strings.Split(filepath.Base(fileName), ".")
	viper.SetConfigName(filepath.Base(splits[0]))
	viper.AddConfigPath(filepath.Dir(fileName))
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}

func checkKey(key string) {
	if !viper.IsSet(key) {
		fmt.Printf("Configuration key %s not found; aborting \n", key)
		os.Exit(1)
	}
}

func MustGetString(key string) string {
	checkKey(key)
	return viper.GetString(key)
}

func MustGetInt(key string) int {
	checkKey(key)
	return viper.GetInt(key)
}

func MustGetBool(key string) bool {
	checkKey(key)
	return viper.GetBool(key)
}

func GetString(key string) string {
	return viper.GetString(key)
}
