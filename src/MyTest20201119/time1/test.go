package main

import (
	"fmt"
	"time"
)

const TIME_LAYOUT = "2006-01-02 15:04:05"

func parseWithLocation(name string, timeStr string) (time.Time, error) {
	locationName := name
	if l, err := time.LoadLocation(locationName); err != nil {
		return time.Time{}, err
	} else {
		// 转成带时区的时间
		lt, _ := time.ParseInLocation(TIME_LAYOUT, timeStr, l)
		// 直接转成相对时间
		fmt.Println(time.Now().In(l).Format(TIME_LAYOUT))
		return lt, nil
	}
}
func testTime() {
	str := time.Now().Format("2006-01-02 15:04:05")
	// 指定时区
	t1, err := parseWithLocation("America/Cordoba", str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t1)

	t2, err := parseWithLocation("Asia/Shanghai", str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t2)

	t3, err := parseWithLocation("Asia/Chongqing", str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t3)
}

func main() {
	testTime()
}
