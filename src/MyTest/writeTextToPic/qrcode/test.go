package main

import "github.com/skip2/go-qrcode"

func main() {
	qrcode.WriteFile("{\"key11\":\"123jkkjkngg你好gfccffg\"}", qrcode.Medium, 256, "./qrcode.png")
}
