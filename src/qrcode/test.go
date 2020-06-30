package main

import "github.com/skip2/go-qrcode"

func main() {
	qrcode.WriteFile("{\"key\":\"123jkkjkngggfccffg\"}", qrcode.Medium, 256, "./qrcode.png")
}
