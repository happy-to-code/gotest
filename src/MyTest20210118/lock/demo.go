package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	count := 0
	for i := 0; i < 500; i++ {
		go func() {
			mutex.Lock()
			count += 1
			mutex.Unlock()
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("the count is : ", count)

}

func count(i int) int {
	i += 1
	return i
}
