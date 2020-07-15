package main

import (
	"github.com/gin-gonic/gin"
)

// 我们定义一个通用的格式化的响应数据
// 在Data字段中采用空接口类型来实际存放我们的业务数据
type restData struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  bool        `json:"status"`
}

func main7() {
	// mock一个http响应数据
	restdata := &restData{"Hello,BGBiao", "", true}
	restdata1 := &restData{map[string]string{"name": "BGBiao", "website": "https://bgbiao.top"}, "", true}

	// 使用Gin框架启动一个http接口服务
	ginObj := gin.Default()
	ginObj.GET("/api/test", func(c *gin.Context) {
		// 我们的handlerFunc中入参是一个Context结构的引用对象c
		// 因此我们可以使用Context中的JSON方法来返回一个json结构的数据
		// 可用的方法有如下几种，我们可以根据实际需求进行选择
		/*
		   IndentedJSON(code int, obj interface{}): 带缩进的json(消耗cpu和mem)
		   SecureJSON(code int, obj interface{}): 安全化json
		   JSONP(code int, obj interface{})
		   JSON(code int, obj interface{}): 序列化为JSON,并写Content-Type:"application/json"头
		*/
		c.JSON(200, restdata)
	})
	ginObj.GET("/api/test1", func(c *gin.Context) {
		c.IndentedJSON(200, restdata1)
	})

	ginObj.Run("localhost:8080")
}
