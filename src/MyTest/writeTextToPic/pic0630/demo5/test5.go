package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"image/color"
	"log"
)

func main() {
	const (
		S = 1024
	)

	im, err := gg.LoadImage("E:\\20.06.16Project\\GoTest\\src\\pic0630\\demo5\\exc.png")
	if err != nil {
		log.Fatal(err)
	}

	width := im.Bounds().Size().X
	height := im.Bounds().Size().Y

	fmt.Println("im.Bounds().Size().X", width)
	fmt.Println("im.Bounds().Size().Y", height)

	dc := gg.NewContext(width, height)
	dc.DrawImage(im, 0, 0)
	//dc.SetRGB(122, 255, 122)
	//dc.Clear()
	//dc.SetRGB(255, 255, 255)
	if err := dc.LoadFontFace("C:\\Windows\\Fonts\\simsun.ttc", 18); err != nil {
		panic(err)
	}
	// 设置字体颜色
	dc.SetColor(color.RGBA{A: 255})

	dc.DrawRoundedRectangle(0, 0, float64(width), float64(height), 0)
	//dc.DrawImage(im, 0, 0)

	//dc.DrawStringAnchored("小明同学", 121+4*9, 20, 0.5, 0.5)
	dc.DrawStringAnchored("小明同学", 157, 20, 0.5, 0.5)
	dc.DrawStringAnchored("18", 121+8, 50, 0.5, 0.5)
	dc.DrawStringAnchored("男", 121+9, 90, 0.5, 0.5)
	dc.DrawStringAnchored("苏州市相城大道火车站1118号", 121+11*9+8+8, 120, 0.5, 0.5)
	dc.DrawStringAnchored("18932156666", 121+8*6, 150, 0.5, 0.5)
	dc.Clip()

	//dc.DrawImage(im, 0, 0)
	dc.SavePNG("out1.png")
}
