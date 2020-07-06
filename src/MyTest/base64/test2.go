package main

import (
	"encoding/hex"
	"fmt"
)

func main() {

	base64 := "dwfsMuAcoMJorTkCnMEso7DEs7GceoxsxEwjrSebP4I="
	hx := hex.EncodeToString([]byte(base64))
	fmt.Println("Original String: ARVIN")
	fmt.Println()
	fmt.Println(base64 + " ==> " + hx)
}
