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
	dx = 500 // 图片的大小 宽度
	dy = 500 // 图片的大小 高度
	// fontFile = "FZFSK.TTF"
	fontFile = "C:\\Windows\\Fonts\\simsun.ttc"
	fontSize = 20 // 字体尺寸
	fontDPI  = 72 // 屏幕每英寸的分辨率
)

func main() {
	// 需要保存的文件
	imgcounter := 2
	imgfile, _ := os.Create(fmt.Sprintf("%03d.png", imgcounter))
	defer imgfile.Close()
	// 新建一个 指定大小的 RGBA位图
	img := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	// 画背景
	/*for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			// 设置某个点的颜色，依次是 RGBA
			img.Set(x, y, color.RGBA{uint8(x), uint8(y), 0, 255})
		}
	}
	*/

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
	c.SetClip(img.Bounds())
	c.SetDst(img)
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
	err = png.Encode(imgfile, img)
	if err != nil {
		log.Println("生成图片出错")
		log.Fatal(err)
	}
}
