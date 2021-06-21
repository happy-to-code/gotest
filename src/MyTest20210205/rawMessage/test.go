package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// 在程序中使用Json时，有时某个字段其结构是根据其他字段（比如有个类型含义的字段）决定的，这个时候在解析时，
// 需要先解析一部分，进行判断后，再解析出合适的Json结构。这时就需要用到Golang Json包的RawMessage这个对象。
//
// json.RawMessage此类型的定义为type RawMessage []byte，它还实现了json.Marshaler和json.Unmarshaler接口，可以json的延迟编解码。
func main() {
	type Color struct {
		Space string
		Point json.RawMessage // delay parsing until we know the color space
	}
	type RGB struct {
		R uint8
		G uint8
		B uint8
	}
	type YCbCr struct {
		Y  uint8
		Cb int8
		Cr int8
	}

	var j = []byte(`[
		{"Space": "YCbCr", "Point": {"Y": 255, "Cb": 0, "Cr": -10}},
		{"Space": "RGB",   "Point": {"R": 98, "G": 218, "B": 255}}
	]`)
	var colors []Color
	err := json.Unmarshal(j, &colors)
	if err != nil {
		log.Fatalln("error:", err)
	}

	for _, c := range colors {
		var dst interface{}
		switch c.Space {
		case "RGB":
			dst = new(RGB)
		case "YCbCr":
			dst = new(YCbCr)
		}
		err := json.Unmarshal(c.Point, dst)
		if err != nil {
			log.Fatalln("error:", err)
		}
		fmt.Println(c.Space, dst)
	}
}
