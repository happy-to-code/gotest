package main

import (
	"encoding/base64"
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

	fmt.Println("===================================")
	ss := "52a056e17566283a110010b1c4022a5b382dac3dfbd4948b21b14f5960f9a1b5"
	decodeString, _ := base64.StdEncoding.DecodeString(ss)

	fmt.Println("====>", decodeString)
	fmt.Println("====>", string(decodeString))
	toString := base64.StdEncoding.EncodeToString(decodeString)
	fmt.Println("==================>", toString)

}
