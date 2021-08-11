package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	go dealHeight(h)
	router := gin.Default()
	router.POST("/events", events)
	router.POST("/channel", channelTest)
	router.Run(":5000")
}

type ChannelDemo struct {
	Num int
}

var h = make(chan int, 10000)

func addNum2Channel(height int) {
	for i := 0; i < height; i++ {
		h <- i
	}
}

func dealHeight(c chan int) {
	for {
		for i := range c {
			fmt.Println("------>dealHeight:", i)
		}
	}
}

func channelTest(c *gin.Context) {
	var channelDemo ChannelDemo
	err := c.BindJSON(&channelDemo)
	if err != nil {
		panic(err)
	}
	fmt.Printf("参数：%d", channelDemo.Num)

	addNum2Channel(channelDemo.Num)
}

func events(c *gin.Context) {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	fmt.Println(string(buf[0:n]))
	resp := map[string]string{"hello": "world"}
	c.JSON(http.StatusOK, resp)
}
