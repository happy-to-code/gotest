package main

import (
	"bytes"
	"fmt"
)

/**
 * @Author zhangyifei
 * @Description  拼接字符串
 * @Date 2020/7/17 16:27
 * @Param
 * @return
 */
func main() {
	var buf bytes.Buffer

	buf.WriteString("hello")
	buf.WriteString(" world")
	fmt.Println("buf.string():", buf.String())
}
