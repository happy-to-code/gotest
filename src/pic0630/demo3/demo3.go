package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

func main() {
	//图片，网上随便找了一张
	imgFile, err := os.Open("E:\\20.06.16Project\\GoTest\\src\\pic0630\\demo3\\bg.jpg")
	if err != nil {
		fmt.Println("打开图片出错")
		fmt.Println(err)
		os.Exit(-1)
	}
	defer imgFile.Close()
	//img, err := jpeg.Decode(imgFile)
	img, _, err := image.Decode(imgFile)
	if err != nil {
		fmt.Println("把图片解码为结构体时出错")
		fmt.Println(img)
		os.Exit(-1)
	}

	//水印,用的是我自己支付宝的二维码
	wmbFile, err := os.Open("E:\\20.06.16Project\\GoTest\\src\\pic0630\\demo3\\alipay.png")
	if err != nil {
		fmt.Println("打开水印图片出错")
		fmt.Println(err)
		os.Exit(-1)
	}
	defer wmbFile.Close()
	wmbImg, err := png.Decode(wmbFile)
	if err != nil {
		fmt.Println("把水印图片解码为结构体时出错")
		fmt.Println(err)
		os.Exit(-1)
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

	fmt.Println("添加水印图片结束请查看")
}
