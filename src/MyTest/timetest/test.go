package main

import (
	"fmt"
	"time"
)

func main() {
	// 2020-07-06 15:09:16:089
	var timeStamp uint64 = 1594019356089

	// 返回time对象
	t := time.Unix(int64(timeStamp/1000), 0)

	// 返回string
	dateStr := t.Format("2006年01月02日 15:04:05")

	fmt.Printf("%-10s %-10T %s\n", "dateStr", dateStr, dateStr)

}
