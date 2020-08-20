package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func main() {
	// 创建或者打开数据库
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 创建表
	err = db.Update(func(tx *bolt.Tx) error {
		// 获取或者新建一个表（bucket）
		b, _ := tx.CreateBucketIfNotExists([]byte("BlockBucket"))

		// 往表里面存储数据
		if b != nil {
			err := b.Put([]byte("1"), []byte("Send 1000 BTC To 冠希哥......"))
			if err != nil {
				log.Panic("数据存储失败......")
			}
		}

		// 返回nil，以便数据库处理相应操作
		return nil
	})
	// 更新失败
	if err != nil {
		log.Panic(err)
	}

	// 查看数据
	err = db.View(func(tx *bolt.Tx) error {
		fmt.Println("------")
		// 获取BlockBucket表对象
		b := tx.Bucket([]byte("BlockBucket"))

		// 往表里面存储数据
		if b != nil {
			data0 := b.Get([]byte("1"))
			fmt.Printf("数据:%s\n", data0)
			data := b.Get([]byte("ll"))
			fmt.Printf("数据:%s\n", data)
		}

		// 返回nil，以便数据库处理相应操作
		return nil
	})
	// 更新失败
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("===================================")
}
