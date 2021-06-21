package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	onchianTime := 1605595386
	timeStr1 := strconv.Itoa(onchianTime)
	fmt.Println("timeStr1:", timeStr1)
	fmt.Println("len:", len(timeStr1))
	if len(timeStr1) == 13 {
		onchianTime = onchianTime / 1000
	}
	// yyyy-MM-dd'T'HH:mm:ssX
	timeStr := time.Unix(int64(onchianTime), 0).Format("2006-01-02'T' 15:04:05X")
	fmt.Println(timeStr)

	timeStr2 := time.Unix(int64(onchianTime), 0).Format("2006-01-02 15:04:05")
	fmt.Println(timeStr2)

	t7, _ := time.ParseInLocation("2006-01-02T15:04:05", "2018-10-01T16:27:00", time.Local)

	fmt.Println(t7)
}
