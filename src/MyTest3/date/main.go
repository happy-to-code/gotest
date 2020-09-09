package main

import (
	"fmt"
	"time"
)

func main() {
	timeStr := "2019-09-16"
	loc, _ := time.LoadLocation("Local")
	location, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)

	if err != nil {
		fmt.Println("err::", err)
	}
	fmt.Println("location:", location.Unix())

}
