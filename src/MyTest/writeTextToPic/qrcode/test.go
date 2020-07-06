package main

import "github.com/skip2/go-qrcode"

func main() {
	qrcode.WriteFile("http://www.baidu.com", qrcode.Highest, 290, "./qrco7de.png")
}
