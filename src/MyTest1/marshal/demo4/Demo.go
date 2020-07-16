package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Girl struct {
	Name       string
	Age        int    `json:"age"`
	Gender     string `json:"gender"`
	Where      string `json:"where"`
	Is_married bool   `json:"is_married"`
}

func main() {
	g := Girl{"satori", 16, "f", "东方地灵殿", false}

	// 创建缓存
	buf := new(bytes.Buffer)
	// 把指针丢进去
	enc := gob.NewEncoder(buf)

	// 调用Encode进行序列化
	if err := enc.Encode(g); err != nil {
		fmt.Println(err)
		return
	} else {
		// 序列化的内容会被放进buf里面
		fmt.Println("buf.String():", buf.String())
	}

	// 	进行反序列化
	// buf2 := new(bytes.Buffer)
	dec := gob.NewDecoder(bytes.NewBuffer(buf.Bytes()))
	var g1 = Girl{}
	dec.Decode(&g1)

	fmt.Println("g1:", g1)

}
