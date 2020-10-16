package main

import (
	"encoding/json"
	"fmt"
)

/**
 * @Author zhangyifei
 * @Description  序列化
 * @Date 2020/7/16 16:12.txt
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
	lesson := Lesson{[]string{"Math", "English", "Chinese"}}
	people := &People{
		Name:   "amy",
		Age:    22,
		Gender: "female",
		Lesson: lesson,
	}
	if b, err := json.Marshal(people); err != nil {
		fmt.Println("Marshal failed...")
	} else {
		fmt.Println("b:", b)
		fmt.Println("string(b):", string(b))
	}
}

// // 打印结果
// [123 34 97 103 101 34 58 50 50 44 34 103 101 110 100 101 114 34 58 34 102 101 109 97 108 101 34 44 34 108 101 115 115 111 110 115 34 58 91 34 77 97 116 104 34 44 34 69 110 103 108 105 115 104 34 44 34 67 104 105 110 101 115 101 34 93 125]
// {"age":22,"gender":"female","lessons["Math","English","Chinese“}
//
// 序列化
