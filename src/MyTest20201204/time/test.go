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
	fmt.Print("---------------------\n")
	var a interface{}
	a = 1
	// num1, ok1 := a.(float64)
	// fmt.Printf("%v%v",num1,ok1)
	if num, ok := a.(int); ok {
		fmt.Println("num:", num)
	} else {
		panic("sss")
	}
}
