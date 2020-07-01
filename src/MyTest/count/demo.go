package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "这是个一455包哈哈含h汉  字和英文的字s符串, This is an apple,123"
	hzc := 0
	num := 0
	word := 0
	space := 0

	for _, v := range s {
		if unicode.Is(unicode.Han, v) {
			hzc++
			continue
		}
		if unicode.Is(unicode.Number, v) {
			num++
			continue
		}
		if 64 < v && v < 122 {
			word++
			continue
		}
		if v == 32 {
			space++
			continue
		}
	}

	fmt.Printf("该字符串中找到中文个数是：%v\n", hzc)
	fmt.Printf("该字符串中找到数字个数是：%v\n", num)
	fmt.Printf("该字符串中找到字母个数是：%v\n", word)
	fmt.Printf("该字符串中找到空格个数是：%v\n", space)

}
