package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	Demo()
}

func Demo() {
	s := "MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAE3a7JNriegwzH3LMK/d0LCGMGkrnS\nxe+QuxttLkUt7uqlHPlOAlywr/ZzMp/Ad510W2a0zUCIhmZky4UQdv3aPA=="
	encodeToString := hex.EncodeToString([]byte(s))
	fmt.Printf("%s", encodeToString)
	fmt.Println("---------------")
	var a uint64 = 10
	// fmt.Println(-a)
	fmt.Println(-int64(a))
}
