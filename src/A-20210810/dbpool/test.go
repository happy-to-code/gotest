package main

import (
	"A-20210810/dbpool/config"
	"A-20210810/dbpool/dao"
	"fmt"
	"gorm.io/gorm"
)

// Product 产品
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	p := new(Product)
	p.Test()
}

// Test 测试
func (Product) Test() {
	config.InitConfig()
	db := dao.GetDB()
	// 自动迁移模式
	// db.AutoMigrate(&Product{})

	// 创建
	// tx := db.Create(&Product{Code: "L1212", Price: 1000})
	exec := db.Exec("insert into product values(?,?)", "L1212", 1000)
	fmt.Println(exec)

}
