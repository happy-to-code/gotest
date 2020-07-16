package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

/**
 * @Author zhangyifei
 * @Description  数据编码和解码
 * @Date 2020/7/16 16:24
 * @Param
 * @return
 */
type People struct {
	Name   string `json:"name"`   // name,小写不导出
	Age    int    `json:"age"`    // age，在 json 字符串中叫 age
	Gender string `json:"gender"` // gender
	Lesson
}

type Lesson struct {
	Lessons []string `json:"lessons"`
}

func main() {

	jsonStr := `{"Age": 18,"name": "Jim" ,"gender": "男","lessons":["English","History"],"Room":201,"n":null,"b":false}`
	strR := strings.NewReader(jsonStr)
	people := &People{}

	// 用 NewDecoder && Decode 进行解码给定义好的结构体对象 people
	err := json.NewDecoder(strR).Decode(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", people)

	// 用 NewEncoder && Encode 把保存的 people 结构体对象编码为 json 保存到文件
	f, err := os.Create("./people.json")
	json.NewEncoder(f).Encode(people)

}
