package main

import (
	"fmt"
	"strings"
)

// 定义一个底层是函数类型的自定义类型
type ProcessBasename func(string) string

func main() {
	pathStr := "a/b/cccc.go"
	fmt.Println(GetBaseName(pathStr, basename))
	fmt.Println("----------------")
	pathStr1 := "a/b/ccfrfcc.go"
	fmt.Println(GetBaseName1(pathStr1, basename))

}

// 移除字符串的路径和 .后缀部分  只取文件名称
func basename(str string) string {
	// 去除路径部分
	slash := strings.LastIndex(str, "/") // 如果没有找到“/” lastIndex 返回-1.json
	str = str[slash+1:]

	// 取出.的前一部分
	if dot := strings.LastIndex(str, "."); dot > 0 {
		str = str[:dot]
	}
	return str
}

//  这个函数使用了函数变量作为参数
func GetBaseName(str string, myBasenameFunc func(string) string) string {
	return myBasenameFunc(str)
}

//  这个函数使用type关键字 自定义的类型（底层是一种函数类型）
func GetBaseName1(str string, processBasename ProcessBasename) string {
	return processBasename(str)
}
