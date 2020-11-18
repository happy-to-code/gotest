package main

import (
	"fmt"
	"strconv"
	"strings"
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

	timeStr := time.Unix(int64(onchianTime), 0).Format("2006-01-02 15:04:05")
	fmt.Println(timeStr)

	fmt.Println(strings.ToLower("Account_Audit_Time"))
}
