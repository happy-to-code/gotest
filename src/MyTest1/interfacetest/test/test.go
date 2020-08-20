package main

import "fmt"

type TestStruct struct{}

func NilOrNot(v interface{}) bool {
	// 这边进行了隐式转换
	// 转换后的变量不仅包含转换前的变量，还包含变量的类型信息 TestStruct
	fmt.Printf("v:%+v\n", v)
	fmt.Println("v==nil:", v == nil)
	return v == nil
}

func main() {
	var s *TestStruct
	fmt.Println("s == nil::::", s == nil)    // #=> true
	fmt.Println("NilOrNot(s):", NilOrNot(s)) // #=> false

}
