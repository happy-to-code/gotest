package main

import (
	"fmt"
	"sort"
)

func main() {
	var studs = Students{
		{"zhangsan", 12, 500},
		{"lisi", 13, 60},
		{"wangwu", 10, 100},
	}
	println(studs.Len())
	println(studs[studs.Len()-1].Name)
	println(studs[0].Name)


	fmt.Println("---------------------")
	// var studs Students
	fmt.Println("students", studs)

	sort.Slice(studs, func(i int, j int) bool {
		return studs[i].Score < studs[j].Score
	})
	fmt.Println("students====>", studs)


	sort.Sort(studs)
	fmt.Println("students", studs)
}

type Student struct {
	Name  string
	Age   int
	Score int
}

type Students []Student

// 根据元素的年龄排序
func (s Students) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

// 交换数据
func (s Students) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// 获取此 slice 的长度
func (s Students) Len() int { return len(s) }
