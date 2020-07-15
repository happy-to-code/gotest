package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
type Context struct {
	// 一个包含size,status和ResponseWriter的结构体
	writermem responseWriter
	// http的请求体(指向原生的http.Request指针)
	Request   *http.Request
	// ResonseWriter接口
	Writer    ResponseWriter

	// 请求参数[]{"Key":"Value"}
	Params   Params
	handlers HandlersChain
	index int8
	// http请求的全路径地址
	fullPath string
	// gin框架的Engine结构体指针
	engine   *Engine
	// 每个请求的context中的唯一键值对
	Keys map[string]interface{}
	// 绑定到所有使用该context的handler/middlewares的错误列表
	Errors errorMsgs
	// 定义了允许的格式被用于内容协商(content)
	Accepted []string
	// queryCache 使用url.ParseQuery来缓存参数查询结果(c.Request.URL.Query())
	queryCache url.Values
	// formCache 使用url.ParseQuery来缓存PostForm包含的表单数据(来自POST,PATCH,PUT请求体参数)
	formCache url.Values
}
*/
/**
 * @Author zhangyifei
 * @Description
 * @Date 2020/7/14 17:42
 * @Param
 * @return
 */
func main() {
	router := gin.Default()

	// router.GET("/someGet", getting)

	// 路由后面加参数，参数必须要传递，不然访问不通过
	router.GET("/someGet/:name", geting)
	router.Run()
}
func geting(c *gin.Context) {
	// c.String(http.StatusOK, "Hello gin")

	// 接收参数
	name := c.Param("name")
	// 使用参数
	c.String(http.StatusOK, "Hello gin %s", name)

}
