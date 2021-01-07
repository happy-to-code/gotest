package main

import (
	"fmt"
	"github.com/gammazero/workerpool"
	"strconv"
)

func main() {
	// wp := workerpool.New(2)
	// requests := []string{"alpha", "beta", "gamma", "delta", "epsilon"}
	//
	// for _, r := range requests {
	// 	r := r
	// 	wp.Submit(func() {
	// 		fmt.Println("Handling request:", r)
	// 	})
	// }
	//
	// wp.StopWait()

	list := makeList()
	fmt.Println("list:", list)
}

func makeList() []string {
	var s []string
	for i := 0; i < 10; i++ {
		wp := workerpool.New(2)
		wp.Submit(func() {
			s = funcName(i, s)
		})

		wp.StopWait()
	}
	return s
}

func funcName(i int, s []string) []string {
	iStr := strconv.Itoa(i)
	s = append(s, iStr)
	return s
}
