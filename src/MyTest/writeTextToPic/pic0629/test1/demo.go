package main

import (
	"github.com/golang/freetype"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"os"
)

func createImage() {
	//图片的宽度
	dx := 200
	//图片的高度
	dy := 200
	imgfile, _ := os.Create("test3.png")
	defer imgfile.Close()

	img := image.NewNRGBA(image.Rect(0, 0, dx, dy))

	//设置每个点的 RGBA (Red,Green,Blue,Alpha(设置透明度))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			//设置一块 白色(255,255,255)不透明的背景
			img.Set(x, y, color.RGBA{255, 255, 255, 255})
		}
	}
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
	f := freetype.NewContext()
	//设置分辨率
	f.SetDPI(72)
	//设置字体
	f.SetFont(font)
	//设置尺寸
	f.SetFontSize(26)
	f.SetClip(img.Bounds())
	//设置输出的图片
	f.SetDst(img)
	//设置字体颜色(红色)
	f.SetSrc(image.NewUniform(color.RGBA{0, 0, 0, 255}))

	//设置字体的位置
	pt := freetype.Pt(40, 40+int(f.PointToFixed(26))>>8)

	_, err = f.DrawString("hello,世界", pt)
	if err != nil {
		log.Fatal(err)
	}

	//以png 格式写入文件
	err = png.Encode(imgfile, img)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	createImage()
}
