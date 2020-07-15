package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/post", postgetting)
	router.Run()
}

func postgetting(c *gin.Context) {
	message := c.PostForm("message")
	status := c.PostForm("status")
	nick := c.DefaultPostForm("nick", "anonymous")
	array := c.PostFormArray("data")
	fmt.Println(array)

	c.JSON(200, gin.H{
		"status":  status,
		"message": message,
		"nick":    nick,
	})
}
