package main

import (
	"fmt"
	"github.com/golang/freetype"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"os"
)

const (
	dx = 800 // 图片的大小 宽度
	dy = 800 // 图片的大小 高度
	// fontFile = "FZFSK.TTF"
	fontFile = "C:\\Windows\\Fonts\\simsun.ttc"
	fontSize = 20 // 字体尺寸
	fontDPI  = 72 // 屏幕每英寸的分辨率
)

func main1() {
	//图片，网上随便找了一张
	imgFile, err := os.Open("E:\\20.06.16Project\\GoTest\\src\\pic0630\\demo3\\bg.jpg")
	if err != nil {
		fmt.Println("打开图片出错")
		fmt.Println(err)
		os.Exit(-1)
	}
	//config, _, _ := image.DecodeConfig(imgFile)
	//fmt.Println("width = ", config.Width)
	//fmt.Println("height = ", config.Height)
	defer imgFile.Close()
	//img, err := jpeg.Decode(imgFile)
	img, _, err := image.Decode(imgFile)
	if err != nil {
		fmt.Println("把图片解码为结构体时出错")
		fmt.Println(img)
		os.Exit(-1)
	}
	//====================================
	// 新建一个 指定大小的 RGBA位图
	imgH := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	// 读字体数据
	fontBytes, err := ioutil.ReadFile(fontFile)
	if err != nil {
		log.Println("读取字体数据出错")
		log.Println(err)
		return
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println("转换字体样式出错")
		log.Println(err)
		return
	}

	c := freetype.NewContext()
	c.SetDPI(fontDPI)
	c.SetFont(font)
	c.SetFontSize(fontSize)
	c.SetClip(imgH.Bounds())
	c.SetDst(imgH)
	c.SetSrc(image.Black)

	fmt.Println("200+int(c.PointToFixed(fontSize)>>8):", 200+int(c.PointToFixed(fontSize)>>8))
	pt := freetype.Pt(200, 200+int(c.PointToFixed(fontSize)>>8)) // 字出现的位置

	_, err = c.DrawString("你好，你知道我是谁吗", pt)
	if err != nil {
		log.Println("向图片写字体出错")
		log.Println(err)
		return
	}
	// 重新设置第二行y的位置
	pt.Y += c.PointToFixed(fontSize)
	pt.X += c.PointToFixed(fontSize)
	_, err = c.DrawString("不知道---", pt)
	if err != nil {
		log.Println("向图片写字体出错")
		log.Println(err)
		return
	}

	// 以PNG格式保存文件
	err = png.Encode(imgFile, imgH)
	if err != nil {
		log.Println("生成图片出错")
		log.Fatal(err)
	}

	/*//水印,用的是我自己支付宝的二维码
	wmbFile, err := os.Open("E:\\20.06.16Project\\GoTest\\src\\pic0630\\demo3\\alipay.png")
	if err != nil {
		fmt.Println("打开水印图片出错")
		fmt.Println(err)
		os.Exit(-1.json)
	}
	defer wmbFile.Close()
	wmbImg, err := png.Decode(wmbFile)
	if err != nil {
		fmt.Println("把水印图片解码为结构体时出错")
		fmt.Println(err)
		os.Exit(-1.json)
	}

	//把水印写在右下角，并向0坐标偏移10个像素
	offset := image.Pt(img.Bounds().Dx()-wmbImg.Bounds().Dx()-10, img.Bounds().Dy()-wmbImg.Bounds().Dy()-10)
	b := img.Bounds()
	//根据b画布的大小新建一个新图像
	m := image.NewRGBA(b)

	//image.ZP代表Point结构体，目标的源点，即(0,0)
	//draw.Src源图像透过遮罩后，替换掉目标图像
	//draw.Over源图像透过遮罩后，覆盖在目标图像上（类似图层）
	draw.Draw(m, b, img, image.ZP, draw.Src)
	draw.Draw(m, wmbImg.Bounds().Add(offset), wmbImg, image.ZP, draw.Over)

	//生成新图片new.jpg,并设置图片质量
	imgw, err := os.Create("new3.jpg")
	jpeg.Encode(imgw, m, &jpeg.Options{100})
	defer imgw.Close()

	fmt.Println("添加水印图片结束请查看")*/
}
