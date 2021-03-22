package main

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

func main() {

	newUUID, _ := uuid.NewUUID()
	println(newUUID.String())
	rand.Seed(time.Now().UnixNano())
	// n := rand.Intn(100)
	n := 90
	fmt.Println("n:", n)

	if n >= 10 && n <= 90 {
		for i := 0; i < 10; i++ {
			n = rand.Intn(100)
			if n >= 10 && n <= 30 {
				fmt.Println("retry:", n, "---", i)
				break
			} else {

				fmt.Println("--->", i, ":", n)
			}
		}
	}

	fmt.Println("--------------last-----------------------------", n)
}
