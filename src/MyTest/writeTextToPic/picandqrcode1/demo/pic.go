package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
	"github.com/skip2/go-qrcode"
	"image"
	"image/color"
	"log"
	"os"
	"unicode"
)

var (
	bgFile    *os.File
	bgImg     image.Image
	qrCodeImg image.Image
	offset    image.Point
)

type Header struct {
	CourtName   string ` json:"courtName"`
	CourtId     int64  ` json:"courtId"`
	BizSystemId string ` json:"bizSystemId"`
	CaseId      string ` json:"caseId"`
	Category    string ` json:"category"`
	SubCategory string ` json:"subCategory"`
	Timestamp   int64  ` json:"timestamp"`
}
type Data struct {
	Sender      string `json:"sender"`
	People      string `json:"people"`
	Idcard      string `json:"idcard"`
	Method      string `json:"method"`
	Address     string `json:"address"`
	SendTime    string `json:"sendTime"`
	SignoffTime string `json:"signoffTime"`
}

type Check struct {
	Digest   string `json:"digest"`
	HashAlgo string `json:"hashAlgo"`
	Sign     string `json:"sign"`
}

type StoreInfo struct {
	Header Header ` json:"header"`
	Data   Data   ` json:"data"`
	Check  Check  ` json:"check"`
}
type PicConfig struct {
	ImageFilePath     string
	BodyFontFilePath  string
	TitleFontFilePath string
	BodyDpi           int64
	QrCode            QrCode
	Text              Text
	Body              Body
	LeftSideText      []LeftSideText
}
type QrCode struct {
	Size    int
	Address string
}
type Text struct {
	MainTitleFontSize   int64
	MainTitleFontColor  []int
	MainTitleContent    string
	MainTitleStartPoint []int
	SubTitleFontSize    int64
	SubTitleFontColor   []int
	SubTitleContent     string
	SubTitleStartPoint  []int
}
type Body struct {
	LeftSideFontSize   int64
	LeftSideFontColor  []int
	bodyTextStartPoint []int
}
type LeftSideText struct {
	Key   string
	Value string
}

func main1() {
	var storeInfoStr = `{
    "check": {
        "digest": "64313237313338663739393632663038643965356663363639353965366562373133346436323432363030376432626230323866383830373265343731306339",
        "hashAlgo": "sha256",
        "sign": "3045022100dabee3a740d23f9ab4672d272bde2795894da1cd3c054856865b4add0d1f1ac10220391171c73e4a6b6872ce671fcbb3dcf4ed29bb87d774625ba1217717d86b1a05"
    },
    "data": {
        "address": "13821324323",
        "signoffTime": "2020-06-11 10:32:19",
        "method": "短信",
        "sender": "相城法院",
        "idcard": "3456784567845678",
        "people": "张三",
        "sendTime": "2020-06-10 15:32:23"
    },
    "header": {
        "bizSystemId": "通达海-电子送达系统",
        "caseId": "XC1001-001",
        "category": "电子送达",
        "courtId": 123,
        "courtName": "相城法院",
        "subCategory": "送达签收",
        "timestamp": 1593772650696
    }
}`
	fmt.Println("storeInfoStr:", storeInfoStr)
	var storeInfo StoreInfo
	var infoByte = []byte(storeInfoStr)
	err := json.Unmarshal(infoByte, &storeInfo)
	if err != nil {
		panic(err)
		return
	}

	fmt.Println("storeInfoStruct::::", storeInfo)

	file := readTemplateFile("delivery.toml")
	fmt.Println("conf.Text.MainTitleFontColor:", file.Text.MainTitleFontColor)
	fmt.Println("conf.Text.MainTitleFontColor:", file.Text.MainTitleFontColor[0])
	fmt.Println("conf.Text.MainTitleFontColor:", file.Text.MainTitleFontColor[1])
	fmt.Println("conf.Text.MainTitleFontColor:", file.Text.MainTitleFontColor[2])

	dc := writeTextToPic(file)

	// bgImg, _ := gg.LoadImage(picPath)

	// 03: 生成二维码
	qrSize := file.QrCode.Size
	content := file.QrCode.Address
	qrCodeImg, err = createAvatar(content, qrSize)
	if err != nil {
		fmt.Println("生成二维码失败:", err)
		return
	}
	dc.DrawImage(qrCodeImg, 2551-458-qrSize, 3437-597-qrSize)
	dc.SavePNG("123467.jpg")

}

//  读取配置文件
func readTemplateFile(templateName string) (conf PicConfig) {
	var templateUrl = "E:/20.06.16Project/GoTest/src/MyTest/writeTextToPic/picandqrcode1/conf/" + templateName

	if _, err := toml.DecodeFile(templateUrl, &conf); err != nil {
		log.Fatal("读取配置文件失败", err)
		return
	}
	return conf
}

/*
	向模板图片中写内容
	templatePicPath 模板图片地址
	fontSize 字体大小，单位像素
*/
func writeTextToPic(conf PicConfig) *gg.Context {
	// 加载模板图片
	templatePicPath := conf.ImageFilePath
	im, err := gg.LoadImage(templatePicPath)
	if err != nil {
		log.Fatal(err)
	}
	// 获取模板图片的尺寸
	width := im.Bounds().Size().X
	height := im.Bounds().Size().Y

	fmt.Println("width:", width)
	fmt.Println("height:", height)

	// dc := gg.NewContext(width, height)
	// dc := gg.NewContext(1513, 2452)
	dc := gg.NewContextForImage(im)

	// template := GetPicTemplate("delivery.json")
	// blocks := template.Blocks
	//
	// // GenerateCert(dc,blocks,_)

	// dc.DrawImage(im, 0, 0)

	// 1、写主标题
	titleFontFilePath := conf.TitleFontFilePath
	mainTitleFontSize := conf.Text.MainTitleFontSize
	mainTitleFontColor := conf.Text.MainTitleFontColor
	mainTitleContent := conf.Text.MainTitleContent
	titleStartPoint := conf.Text.MainTitleStartPoint
	if err := dc.LoadFontFace(titleFontFilePath, float64(mainTitleFontSize)); err != nil {
		panic(err)
	}
	// 设置字体颜色
	dc.SetColor(color.RGBA{R: uint8(mainTitleFontColor[0]), G: uint8(mainTitleFontColor[1]), B: uint8(mainTitleFontColor[2]), A: 255})

	dc.DrawRoundedRectangle(0, 0, float64(width), float64(height), 0)

	titleCount := countNum(mainTitleContent)
	dc.DrawStringAnchored(mainTitleContent, float64(titleStartPoint[0])+titleCount.word*float64(mainTitleFontSize)/4+titleCount.num*float64(mainTitleFontSize)/4, float64(titleStartPoint[1]), 0.5, 0.5)

	// 2、写入副标题
	bodyFontFilePath := conf.BodyFontFilePath
	subTitleFontSize := conf.Text.SubTitleFontSize
	subTitleFontColor := conf.Text.SubTitleFontColor
	subTitleContent := conf.Text.SubTitleContent
	subTitleStartPoint := conf.Text.SubTitleStartPoint
	if err := dc.LoadFontFace(bodyFontFilePath, float64(subTitleFontSize)); err != nil {
		panic(err)
	}
	// 设置字体颜色
	dc.SetColor(color.RGBA{R: uint8(subTitleFontColor[0]), G: uint8(subTitleFontColor[1]), B: uint8(subTitleFontColor[2]), A: 255})
	subTitleCount := countNum(subTitleContent)
	dc.DrawStringAnchored(subTitleContent, float64(subTitleStartPoint[0])+subTitleCount.word*float64(subTitleFontSize)/4+subTitleCount.num*float64(subTitleFontSize)/4, float64(subTitleStartPoint[1]), 0.5, 0.5)

	// 发送人
	sender := "相城法院"
	// 当事人
	people := "张三"
	// 身份证号
	idcard := "345678199303264569"
	// 送达时间
	time := "2020年06月11日 10:32:19"
	// 上链时间
	signoffTime := "2020年06月11日 10:32:26"
	//  送达方式
	category := "短信"

	bh := "TJ-XSSDFD123"

	var fontSize float64 = 90

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

	timeCount := countNum(time)
	dc.DrawStringAnchored(time, 900+timeCount.hzc*fontSize/2+timeCount.num*fontSize/4+60, 1779+10, 0.5, 0.5)

	dc.DrawStringAnchored("同济区块链", 900+5*fontSize/2, 1959+10, 0.5, 0.5)

	signoffTimeCount := countNum(signoffTime)
	dc.DrawStringAnchored(signoffTime, 900+signoffTimeCount.hzc*fontSize/2+signoffTimeCount.num*fontSize/4+60, 2139+10, 0.5, 0.5)

	dc.Clip()

	return dc
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
