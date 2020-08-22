package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type EnterpriseInfo struct {
	ContractAddress             string `binding:"required"` // 合约地址
	EquityCode                  string `binding:"required"` // 股权代码
	EquityAbbreviation          string `binding:"required"` // 股权简称
	EquityClass                 string `binding:"required"` // 股权类别 N版 E版
	TotalEquity                 uint64 `binding:"required"` // 总股本
	FullNameOfCompany           string `binding:"required"` // 公司全称
	EnglishName                 string `binding:"required"` // 公司英文名称
	TypeOfEnterpriseCertificate string `binding:"required"` // 企业证件类别
	EnterpriseCertificateNumber string `binding:"required"` // 企业证件编号
	RegisteredCurrency          string `binding:"required"` // 注册币种
	Extend                      string // 扩展字段，必须是Json串
}

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
