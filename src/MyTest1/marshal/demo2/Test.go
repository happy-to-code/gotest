package main

import (
	"encoding/json"
	"fmt"
)

/**
 * @Author zhangyifei
 * @Description  序列化-->传输-->反序列化
 * @Date 2020/7/16 16:16
 * @Param
 * @return
 */
type Student struct {
	Name    string
	Age     int
	Guake   bool
	Classes []string
	Price   float32
}

func (s *Student) ShowStu() {
	fmt.Println("show Student :")
	fmt.Println("\tName\t:", s.Name)
	fmt.Println("\tAge\t:", s.Age)
	fmt.Println("\tGuake\t:", s.Guake)
	fmt.Println("\tPrice\t:", s.Price)
	fmt.Printf("\tClasses\t: ")
	for _, a := range s.Classes {
		fmt.Printf("%s ", a)
	}
	fmt.Println("")
}

func main() {
	st := &Student{
		"XiaoMing",
		16,
		true,
		[]string{"Math", "English", "Chinese"},
		9.99,
	}
	fmt.Println("before JSON encoding :")
	st.ShowStu()

	b, err := json.Marshal(st)
	if err != nil {
		fmt.Println("encoding failed")
	} else {
		fmt.Println("encoded data : ")
		fmt.Println("b:", b)
		fmt.Println("string(b):", string(b))
	}

	ch := make(chan string, 1)
	go func(c chan string, str string) {
		c <- str
	}(ch, string(b))
	strData := <-ch

	fmt.Println("--------------------------------")
	stb := &Student{}
	stb.ShowStu()
	err = json.Unmarshal([]byte(strData), &stb)
	if err != nil {
		fmt.Println("Unmarshal failed")
	} else {
		fmt.Println("Unmarshal success")
		stb.ShowStu()
	}
}
