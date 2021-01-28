package main

import (
	"MyTest20201204/cache"
	"fmt"
	"strings"
)

func main() {
	newCache := cache.NewCache(30, 30)

	newCache.Set("Key", "v1", 30)

	value, b := newCache.Get("Key")
	fmt.Println(value, b)
	fmt.Println(strings.Contains("errorcd","error"))

}
