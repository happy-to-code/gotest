package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main2() {
	router := gin.Default()

	router.GET("/someGet", getting2)
	router.Run()
}
// 127.0.0.1:8080/someGet?firstname=yida&lastname=zhang
// 127.0.0.1:8080/someGet?lastname=zhang
func getting2(c *gin.Context) {
	// 从URL中获取值
	// 如果从URL中为获取到firstname  name就用Guest替换
	firstname := c.DefaultQuery("firstname", "Guest")

	lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}
