package main

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

var engine *xorm.Engine

func main() {
	var err error
	engine, err = xorm.NewEngine("sqlite3", "./test.db")
	fmt.Println("err:", err)
}
