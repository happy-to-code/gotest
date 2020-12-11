package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取本地location
	toBeCharge := "2020/12/01'T' 10:43:00X"                           // 待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	timeLayout := "2006/01/02'T' 15:04:05X"                           // 转化所需模板
	loc, _ := time.LoadLocation("Local")                              // 重要：获取时区
	theTime, err := time.ParseInLocation(timeLayout, toBeCharge, loc) // 使用模板在对应时区转化为time.time类型
	if err != nil {
		fmt.Printf("err:%+v\n", err)
	}
	fmt.Println("theTime:", theTime) // 打印输出theTime 2015-01-01 15:15:00 +0800 CST
	sr := theTime.Unix()             // 转化为时间戳 类型是int64
	fmt.Println(sr)                  // 打印输出时间戳 1420041600

	// // 时间戳转日期
	// dataTimeStr := time.Unix(sr, 0).Format(timeLayout) // 设置时间戳 使用模板格式化为日期字符串
	// fmt.Println(dataTimeStr)

	// s:=toBeCharge+"/"+strconv.Itoa(1)

}
