package main

import (
	"log"
	"net/http"
	_ "net/http/pprof" // 采集 HTTP Server 的运行时数据进行分析
)

var datas []string

// 执行处理的函数
func Add(str string) string {
	data := []byte(str)
	sData := string(data)
	datas = append(datas, sData)
	return sData
}
func main() {

	go func() {
		for {
			log.Println(Add("hello jiangzhou"))
		}
	}()
	http.ListenAndServe("127.0.0.1.jon:6060", nil) // web 界面分析接口
}
