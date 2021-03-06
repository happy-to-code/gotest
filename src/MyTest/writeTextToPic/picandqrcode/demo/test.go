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
	"reflect"
	"unicode"
)

var (
	bgFile    *os.File
	bgImg     image.Image
	qrCodeImg image.Image
	offset    image.Point
)

func main() {

	avatar, _ := createAvatar("http://www.baidu.com", 290)
	i, _ := os.Create(path.Base("二维码.png"))
	png.Encode(i, avatar)

}

func main2() {
	// 模板图片路径
	var templatePicPath = "E:\\20.06.16Project\\GoTest\\src\\MyTest\\writeTextToPic\\picandqrcode\\证书222.jpg"
	// 字体大小
	var fontSize float64 = 60
	// 二维码内容
	var content = "==1111hello,world 你好1233333 ？"
	// 二维码尺寸大小
	var qrSize = 290

	picPath := writeTextToPic(templatePicPath, fontSize)
	fmt.Println("picPath:", picPath)

	bgImg, _ := gg.LoadImage(picPath)

	// 03: 生成二维码
	qrCodeImg, err = createAvatar(content, qrSize)
	if err != nil {
		fmt.Println("生成二维码失败:", err)
		return
	}

	offset = image.Pt(bgImg.Bounds().Size().X-458-qrSize, bgImg.Bounds().Size().Y-597-qrSize)
	b := bgImg.Bounds()
	m := image.NewRGBA(b)
	// Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point, op Op)
	draw.Draw(m, b, bgImg, image.Point{X: 0, Y: 0}, draw.Src)
	draw.Draw(m, qrCodeImg.Bounds().Add(offset), qrCodeImg, image.Point{X: 0, Y: 0}, draw.Src)

	fmt.Println("type:", reflect.TypeOf(m))
	// 创建图片
	i, _ := os.Create(path.Base("11111.png"))
	_ = png.Encode(i, m)

	defer i.Close()
	defer os.Remove(picPath)
}

func main1() {
	// 模板图片路径
	var templatePicPath = "E:\\20.06.16Project\\GoTest\\src\\MyTest\\writeTextToPic\\picandqrcode\\testPic.png"
	// 字体大小
	var fontSize float64 = 18

	writeTextToPic(templatePicPath, fontSize)
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

	fmt.Println("width:", width)
	fmt.Println("height:", height)

	dc := gg.NewContext(width, height)
	dc.DrawImage(im, 0, 0)

	if err := dc.LoadFontFace("E:\\20.06.16Project\\GoTest\\src\\MyTest\\writeTextToPic\\picandqrcode\\simsun.ttc", fontSize); err != nil {
		panic(err)
	}
	// 设置字体颜色
	dc.SetColor(color.RGBA{201, 44, 46, 255})

	dc.DrawRoundedRectangle(0, 0, float64(width), float64(height), 0)

	// dc.DrawStringAnchored("小明同学", float64(231-4+utf8.RuneCountInString("小明同学")*9), 350, 0.5, 0.5)
	// dc.DrawStringAnchored("18", 236, 370, 0.5, 0.5)
	// dc.DrawStringAnchored("男", 236, 390, 0.5, 0.5)
	// dc.DrawStringAnchored("苏州市相城大道火车站1118号", 344, 410, 0.5, 0.5)
	// dc.DrawStringAnchored("18932156666", 276.5, 430, 0.5, 0.5)

	// 发送人
	sender := "相城法院"
	// 当事人
	people := "张三"
	// 身份证号
	idcard := "345678199303264569"
	// 送达时间
	time := "2020-06-11 10:32:19"
	// 上链时间
	signoffTime := "1593336799366"
	// signoffTimeInt64, err := strconv.ParseInt(signoffTime, 10, 64)
	// timetest.Unix(int64(block.Header.Timestamp), 0).Format("2006-01-02 15:04:05")
	// timetest.Unix(signoffTimeInt64, 0).Format("2006-01-02 15:04:05")
	//  送达方式
	category := "短信"

	bh := "TJ-XSSDFD123"

	bhCount := countNum(bh)
	dc.DrawStringAnchored(bh, 1140+bhCount.word*fontSize/4+bhCount.num*fontSize/4, 780, 0.5, 0.5)

	// 设置字体颜色
	dc.SetColor(color.RGBA{A: 255})

	senderCount := countNum(sender)
	dc.DrawStringAnchored(sender, 900+senderCount.hzc*fontSize/2+senderCount.num*fontSize/4, 679+360+20, 0.5, 0.5)

	peopleCount := countNum(people)
	dc.DrawStringAnchored(people, 900+peopleCount.hzc*fontSize/2+peopleCount.num*fontSize/4, 1039+180+10+10, 0.5, 0.5)

	idcardCount := countNum(idcard)
	dc.DrawStringAnchored(idcard, 900+idcardCount.word*fontSize/4+idcardCount.num*fontSize/4, 1419+10, 0.5, 0.5)

	categoryCount := countNum(category)
	dc.DrawStringAnchored(category, 900+categoryCount.hzc*fontSize/2+categoryCount.num*fontSize/4, 1599+10, 0.5, 0.5)

	// timeCount := countNum(timetest)
	dc.DrawStringAnchored(time, 900+19*fontSize/4, 1779+10, 0.5, 0.5)

	dc.DrawStringAnchored("同济区块链", 900+5*fontSize/2, 1959+10, 0.5, 0.5)

	signoffTimeCount := countNum(signoffTime)
	dc.DrawStringAnchored(signoffTime, 900+signoffTimeCount.hzc*fontSize/2+signoffTimeCount.num*fontSize/4, 2139+10, 0.5, 0.5)

	dc.Clip()

	// 使用uuid生成随机字符串---》用于作为图片的名字
	u, err := uuid.NewV1()
	if err != nil {
		log.Println("uuid生成出错!")
	}
	picName := u.String() + ".png"
	// picName := "123" + ".png"
	dc.SavePNG(picName)

	// 获取当前项目的路径
	currentPath := GetCurrentPath()
	// 项目路径+图片名称
	picPath = currentPath + "\\" + picName
	return picPath
}

var err error

// 创建二维码
func createAvatar(content string, qrCodeSize int) (image.Image, error) {
	if len(content) == 0 {
		log.Fatal("要生成的二维码内容为空")
		return nil, nil
	}
	if qrCodeSize <= 0 {
		qrCodeSize = 150
	}
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

// 获取当前项目路径
func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

type CountInfo struct {
	hzc   float64
	num   float64
	word  float64
	space float64
}

// 统计汉字 、空格等个数
func countNum(str string) (count CountInfo) {
	var countInfo CountInfo

	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			countInfo.hzc++
			continue
		}
		if unicode.Is(unicode.Number, v) {
			countInfo.num++
			continue
		}
		if 64 < v && v < 122 {
			countInfo.word++
			continue
		}
		if v == 32 {
			countInfo.space++
			continue
		}
	}
	return countInfo
}
