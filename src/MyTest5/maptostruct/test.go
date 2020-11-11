package main

import (
	"fmt"
	"github.com/goinggo/mapstructure"
)

type Person struct {
	Name string
	Age  int
	Eye  []string
}

func MapToStruct() {
	mapInstance := make(map[string]interface{})
	mapInstance["Name"] = "liang"
	mapInstance["Age"] = 28
	mapInstance["Eye"] = []string{"eye1", "eye2"}

	var person Person
	// 将 map 转换为指定的结构体
	if err := mapstructure.Decode(mapInstance, &person); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("map2struct后得到的 struct 内容为:%v\n", person)
	fmt.Println(person.Age, person.Name)
}

func main() {
	MapToStruct()
}
