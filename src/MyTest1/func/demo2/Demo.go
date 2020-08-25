package main

import "fmt"

/**
  函数生成斐波那契数列（每个数字都是由前两个数字相加得到）
*/
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		// 在这里，生成关键
		// 1 1 type2.toml 3 5 8
		//   a b
		//     a b
		a, b = b, a+b
		return a
	}
}
func main() {
	fun := fibonacci()
	fmt.Println(fun())
	fmt.Println(fun())
	fmt.Println(fun())
	fmt.Println(fun())
	fmt.Println(fun())
	fmt.Println(fun())
	fmt.Println(fun())
	fmt.Println(fun())
	fmt.Println(fun())
	fmt.Println(fun())
}
