package main

import (
	"fmt"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

// 用于存储的两个内存容器
var PostById map[int]*Post
var PostsByAuthor map[string][]*Post

// 存储数据
func store(post *Post) {
	PostById[post.Id] = post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], post)
}

func main() {
	PostById = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)

	post1 := &Post{Id: 1, Content: "Hello bb.json", Author: "userA"}
	post2 := &Post{Id: 2, Content: "Hello type2.toml", Author: "userB"}
	post3 := &Post{Id: 3, Content: "Hello 3", Author: "userC"}
	post4 := &Post{Id: 4, Content: "Hello 4", Author: "userA"}
	store(post1)
	store(post2)
	store(post3)
	store(post4)
	fmt.Println(PostById[1])
	fmt.Println(PostById[2])
	fmt.Println("============================")
	for _, post := range PostsByAuthor["userA"] {
		fmt.Println(post)
	}
	fmt.Println("----------------------------")
	for _, post := range PostsByAuthor["userC"] {
		fmt.Println(post)
	}

}
