package main

import (
	"encoding/hex"
	"fmt"
	"time"
)

func main() {
	curNow := time.Now().Local()
	fmt.Printf("curNow:%+v\n", curNow)

	hex.EncodeToString()
}
