package main

import (
	"encoding/json"
	"fmt"
)

/**
 * @Author zhangyifei
 * @Description  反序列化
 * @Date 2020/7/16 16:12
 * @Param
 * @return
 */
type People struct {
	Name   string `json:"name"`   // name,小写不导出
	Age    int    `json:"age"`    // age
	Gender string `json:"gender"` // gender
	Lesson
}

type Lesson struct {
	Lessons []string `json:"lessons"`
}

func main() {
	jsonstr := `{"Age": 18,"name": "Jim" ,"gender": "男","lessons":["English","History"],"Room":201,"n":null,"b":false}`

	// 反序列化
	var people People
	if err := json.Unmarshal([]byte(jsonstr), &people); err == nil {
		fmt.Println("struct people:")
		fmt.Println(people)
	}

	// 反序列化 json 字符串中的一部分
	var lessons Lesson
	if err := json.Unmarshal([]byte(jsonstr), &lessons); err == nil {
		fmt.Println("struct lesson:")
		fmt.Println(lessons)
	}

	// 反序列化 json 字符串数组
	jsonstr = `["English","History"]`
	var str []string
	if err := json.Unmarshal([]byte(jsonstr), &str); err == nil {
		fmt.Println("struct str：")
		fmt.Println(str)
	}
}

// 打印结果
// 　　struct people:
// 　　{ 18 男 {[English History]}}
// 　　struct lesson:
// 　　{[English History]}
// 　　struct str：
// 　　[English History]
//
// 反序列化
