package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"strings"
	"time"
)

func main() {
	m := make(map[string]string)
	t0 := time.Now()
	for i := 0; i < 10000; i++ {
		uuid := createUUID()
		m[uuid] = uuid
	}
	t1 := time.Now()
	fmt.Println("time:", t1)
	fmt.Println("time:", t1.Unix()-t0.Unix())
	fmt.Println("length:", len(m))
}

// 生成UUID
func createUUID() string {
	ul := uuid.NewV4()
	return strings.Replace(ul.String(), "-", "", -1)
}
