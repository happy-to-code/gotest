package main

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
	"io/ioutil"
)

func main() {
	// 读取json-schema文件文件
	jsonSchema, err := ioutil.ReadFile("./json-schema.json")
	if err != nil {
		panic(err)
	}
	sourceDataModel := ""
	schemaLoader := gojsonschema.NewBytesLoader(jsonSchema)                // json-schema lodaer
	documentLoader := gojsonschema.NewBytesLoader([]byte(sourceDataModel)) // 格式化转换
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)     // 校验
	if err != nil {
		panic(err)
	}
	msg := ""
	if !result.Valid() {
		for _, desc := range result.Errors() {
			msg = msg + desc.String() + "\n"
		}
		fmt.Println(msg)
		return
	}
	fmt.Println("success")
}
