package main

import (
	"fmt"
	"github.com/satori/go.uuid"
)

func main() {
	// 创建
	u1 := uuid.NewV4()
	fmt.Printf("UUIDv4: %s\n", u1)
	fmt.Println("UUIDStr:", u1.String())

	v1 := uuid.NewV1()
	fmt.Println("UUIDV1Str:", v1.String())
	picName := v1.String() + ".png"
	fmt.Println("=====:", picName)
	// os.Remove("E:\\20.06.16Project\\GoTest\\90daeacd-bb6d-11ea-a501-40b034477939.png")

	// // 解析
	// u2, err := uuid.FromString("f5394eef-e576-4709-9e4b-a7c231bd34a4")
	// if err != nil {
	// 	fmt.Printf("Something gone wrong: %s", err)
	// 	return
	// }
	// fmt.Printf("Successfully parsed: %s", u2)
}
