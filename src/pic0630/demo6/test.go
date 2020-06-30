package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/golang/freetype"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	im, err := gg.LoadImage("E:\\20.06.16Project\\GoTest\\src\\pic0630\\demo5\\exc.png")
	if err != nil {
		log.Fatal(err)
	}

	img := image.NewRGBA(image.Rect(0, 0, 320, 240))

	x, y := 100, 100
	addLabel(img, x, y, "Test123")
	png.Encode(os.Stdout, img)
}

func addLabel(img *image.RGBA, x, y int, label string) {
	//读取字体数据
	fontBytes, err := ioutil.ReadFile("C:\\Windows\\Fonts\\simsun.ttc")
	if err != nil {
		log.Println(err)
	}
	//载入字体数据
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println("load front fail", err)
	}
	c := freetype.NewContext()
	//设置字体
	c.SetFont(font)
	c.SetDst(img)
	size := 12.0 // font size in pixels
	pt := freetype.Pt(x, y+int(c.PointToFixed(size)>>6))

	if _, err := c.DrawString(label, pt); err != nil {
		fmt.Println(err)
	}
}
