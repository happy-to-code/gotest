package main

import (
	"encoding/json"
	"fmt"
)

type Class struct {
	ClassName  string
	StudentNum int
	Student    Student
}
type Student struct {
	Name string
	Age  int
	Book Book
}
type Book struct {
	BookName string
	BookNum  int
}

func main() {
	var b = Book{"语文", 3}
	fmt.Println("b", b)
	var mm = make(map[string]interface{})
	mm["BookName"] = "数学"
	mm["BookNum"] = 2
	for k, v := range mm {
		fmt.Println(k, "===", v)
		if k == "BookName" {
			b.BookName = fmt.Sprintf("%v", v)
			continue
		}
		if k == "BookNum" {
			b.BookNum = v.(int)
			continue
		}
	}
	fmt.Println("b2", b)
	fmt.Println("==========================")
	c := Class{
		ClassName:  "2班",
		StudentNum: 2,
		Student: Student{
			Name: "小明",
			Age:  16,
			Book: Book{
				BookName: "语文",
				BookNum:  28,
			},
		},
	}

	fmt.Printf("c %+v\n", c)
	var cmap map[string]interface{}
	marshal, _ := json.Marshal(c)
	json.Unmarshal(marshal, &cmap)
	fmt.Printf("cmap %+v\n", cmap)

	for k, v := range cmap {
		fmt.Println(k, "\t", v)

	}

	m := make(map[string]interface{})
	m["Name"] = "小华"

}
