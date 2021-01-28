package main

import "fmt"

func main() {
	td()
}

func td() {
	/* 定义局部变量 */
	var grade string
	var marks int = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
		// default:
		// 	grade = "D"
	}

	fmt.Printf("你的等级是 %s\n", grade)
}
