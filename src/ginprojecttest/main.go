package main

import (
	orm "ginprojecttest/database"
	"ginprojecttest/router"
)

func main() {
	// 当整个项目流程关闭后就关闭数据库链接
	defer orm.Db.Close()

	r := router.InnitRouter()
	r.Run(":8081")
}
