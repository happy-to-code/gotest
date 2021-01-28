package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestCache_Set(t *testing.T) {
	InitCache()

	nonceCache.Set("kkkk", "123", time.Second*5)

	for true {
		get, b := nonceCache.Get("kkkk")
		if b {
			fmt.Printf("%v%v\n", b, get.(string))
		} else {
			fmt.Println("=================")
		}
	}

}
