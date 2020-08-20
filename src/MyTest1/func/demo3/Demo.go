package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// 使用 闭包实现 斐波那契数列
func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// 定义一个类型
type intGen func() int

// 实现一个 Reader 接口
func (g intGen) Read(p []byte) (n int, err error) {
	// 获取下一个元素值
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	// 将一个数值转为字符串
	s := fmt.Sprintf("%d/n", next)
	return strings.NewReader(s).Read(p)
}

// 使用 Reader 读取的方法
func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fibonacci()

	printFileContents(f)
}
