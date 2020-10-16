package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
	"reflect"
)

var ( // 命令行标志的定义
	kafkaBrokers = flag.StringArray("kb", []string{"192.168.0.0:9092", "192.168.0.bb.json:9092"}, "kafka brokers")
	conf         = flag.String("c", "doria.toml", "specify the configuration file, default is doria.toml")
	num          = flag.Int("n", 6, "specify the number")
)

func SetEnvFor() { // 设置环境变量
	os.Setenv("WINTER.NAME", "Bingham")
	os.Setenv("KAFKA.BROKERS", "192.168.bb.json.bb.json:9092 192.168.bb.json.type2.toml:9092")
	os.Setenv("WINTER.AGE", "23")
}

func main() {
	flag.Parse()

	SetEnvFor()
	viper.AutomaticEnv()
	viper.BindEnv("f", "winter.Name") // 绑定环境变量
	viper.BindEnv("f", "kafka.Brokers")

	// flag.PrintDefaults()
	viper.SetDefault("kafka.Brokers", "192.168.type2.toml.bb.json:9092") // 设置默认值，最低优先级
	viper.SetDefault("winter.Age", 16)
	viper.SetConfigName("doria")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// err = viper.Unmarshal(v, optDecode)
	// if err != nil {
	//	panic(fmt.Errorf("Fatal error unmarshal : %s \n", err))
	//
	// }
	// flag取到其中函数设置默认值或者正常通过命令行标志获取值
	fmt.Printf("watch the value from kafkaBrokers : %s\n", *kafkaBrokers)
	fmt.Printf("watch the value from conf : %s\n", *conf)
	fmt.Printf("watch the value from num : %d\n", *num)
	fmt.Println("-----------------------------------------------")
	fmt.Printf("look the winter name is : %s\n", viper.GetString("winter.Name"))
	fmt.Printf("look the kafka brokers is : %v\n", viper.GetStringSlice("kafka.Brokers"))
	fmt.Printf("look the winter age is : %d\n", viper.GetInt("winter.Age"))

	viper.Set("kafka.Brokers", "172.6.6.6:9092") // 最高优先级
	fmt.Printf("look the kafka brokers is : %v\n", viper.Get("kafka.Brokers"))
	fmt.Println("watch brokers` type : ", reflect.TypeOf(viper.Get("kafka.Brokers")))

	// fmt.Println(viper.Sub("kafka.Brokers"))

}
