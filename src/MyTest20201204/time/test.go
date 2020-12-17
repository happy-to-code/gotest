package main

import (
	"fmt"
	"time"
)

func main() {
	loc, _ := time.LoadLocation("Local")

	fundClosingTime, _ := time.ParseInLocation("2006-01-02", "2020-12-16 14:27:00", loc)
	fmt.Println(fundClosingTime.Unix())
	fmt.Println(fundClosingTime)
}
