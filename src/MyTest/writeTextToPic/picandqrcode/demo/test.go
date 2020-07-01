package main

import (
	"errors"
	"fmt"
	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
	uuid "github.com/satori/go.uuid"
	"github.com/skip2/go-qrcode"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"path"
)

var (
	bgFile    *os.File
	bgImg     image.Image
	qrCodeImg image.Image
	offset    image.Point
)

func main() {
	var templatePicPath = "E:\\20.06.16Project\\GoTest\\src\\MyTest\\writeTextToPic\\picandqrcode\\testPic.png"
	var fontSize float64 = 18
	picPath := writeTextToPic(templatePicPath, fontSize)
	fmt.Println("picPath:", picPath)

	bgImg, _ := gg.LoadImage(picPath)

	// 03: 生成二维码
	var content = "hello,world 你好 ？"
	var qrSize = 150
	qrCodeImg, err = createAvatar(content, qrSize)
	if err != nil {
		fmt.Println("生成二维码失败:", err)
		return
	}

	offset = image.Pt(426, 475)
	b := bgImg.Bounds()
	m := image.NewRGBA(b)
	// Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point, op Op)
	draw.Draw(m, b, bgImg, image.Point{X: 0, Y: 0}, draw.Src)
	draw.Draw(m, qrCodeImg.Bounds().Add(offset), qrCodeImg, image.Point{X: 0, Y: 0}, draw.Over)

	// 创建图片
	i, _ := os.Create(path.Base("a778899.png"))
	_ = png.Encode(i, m)

	defer i.Close()
	defer os.Remove(picPath)
}

/*
	向模板图片中写内容
	templatePicPath 模板图片地址
	fontSize 字体大小，单位像素
*/
func writeTextToPic(templatePicPath string, fontSize float64) (picPath string) {
	// 加载模板图片
	im, err := gg.LoadImage(templatePicPath)
	if err != nil {
		log.Fatal(err)
	}
	// 获取模板图片的尺寸
	width := im.Bounds().Size().X
	height := im.Bounds().Size().Y

	dc := gg.NewContext(width, height)
	dc.DrawImage(im, 0, 0)

	if err := dc.LoadFontFace("C:\\Windows\\Fonts\\simsun.ttc", fontSize); err != nil {
		panic(err)
	}
	// 设置字体颜色
	dc.SetColor(color.RGBA{A: 255})

	dc.DrawRoundedRectangle(0, 0, float64(width), float64(height), 0)
	// dc.DrawImage(im, 0, 0)

	// dc.DrawStringAnchored("小明同学", 121+4*9, 20, 0.5, 0.5)
	dc.DrawStringAnchored("小明同学", 157, 20, 0.5, 0.5)
	dc.DrawStringAnchored("18", 121+8, 50, 0.5, 0.5)
	dc.DrawStringAnchored("男", 121+9, 90, 0.5, 0.5)
	dc.DrawStringAnchored("苏州市相城大道火车站1118号", 121+11*9+8+8, 120, 0.5, 0.5)
	dc.DrawStringAnchored("18932156666", 121+8*6, 150, 0.5, 0.5)
	dc.Clip()

	// 使用uuid生成随机字符串---》用于作为图片的名字
	u, err := uuid.NewV1()
	if err != nil {
		log.Println("uuid生成出错!")
	}
	picName := u.String() + ".png"
	dc.SavePNG(picName)

	// 获取当前项目的路径
	currentPath := GetCurrentPath()

	picPath = currentPath + "\\" + picName
	return picPath
}

var err error

// 创建二维码
func createAvatar(content string, qrCodeSize int) (image.Image, error) {
	var (
		bgImg image.Image
	)

	bgImg, err = createQrCode(content, qrCodeSize)
	return bgImg, err
}

func createQrCode(content string, qrCodeSize int) (img image.Image, err error) {
	var qrCode *qrcode.QRCode
	qrCode, err = qrcode.New(content, qrcode.Highest)

	if err != nil {
		return nil, errors.New("创建二维码失败")
	}
	qrCode.DisableBorder = true
	img = qrCode.Image(qrCodeSize)
	return img, nil
}

func ImageResize(src image.Image, w, h int) image.Image {
	return resize.Resize(uint(w), uint(h), src, resize.Lanczos3)
}

// 获取当前项目路径，比如：E:/abc/data/test
func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
