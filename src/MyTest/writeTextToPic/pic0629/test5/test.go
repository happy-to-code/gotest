package main

import (
	"github.com/fogleman/gg"
	"log"
)

func main() {
	const S = 1024
	im, err := gg.LoadImage("E:\\20.06.16Project\\GoTest\\src\\pic0629\\test5\\target.png")
	if err != nil {
		log.Fatal(err)
	}

	dc := gg.NewContext(949, 745)
	dc.SetRGB(255, 18, 61)
	dc.Clear()
	dc.SetRGB(255, 255, 255)
	if err := dc.LoadFontFace("C:\\Windows\\Fonts\\simsun.ttc", 96); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored("==Hello, world!你好！==", S/2, S/2, 0.5, 0.5)

	dc.DrawRoundedRectangle(0, 0, 949, 745, 0)
	dc.DrawImage(im, 0, 0)
	dc.DrawStringAnchored("Hello, world!", S/2, S/2, 0.5, 0.5)
	dc.Clip()
	dc.SavePNG("out11.png")
}
