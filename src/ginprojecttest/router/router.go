package router

import (
	. "ginprojecttest/api"
	"github.com/gin-gonic/gin"
)

/*
InnitRouter 路由初始化
*/
func InnitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/users", Users) // 查询的接口是 users , 复数
	router.POST("/user", Store)
	router.PUT("/user/:id", Update)
	router.DELETE("/user/:id", Del)
	return router
}
