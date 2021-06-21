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
	// "image/jpeg"
	"image/png"
	"log"
	"os"
	"path"
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

func main2() {
	qrSize := 320
	content := "file.QrCode.Address"
	qrCodeImg, err = createAvatar(content, qrSize)
	if err != nil {
		fmt.Println("生成二维码失败:", err)
		return
	}
	i, _ := os.Create(path.Base("二维码1212.png"))
	png.Encode(i, qrCodeImg)
}
func main9() {
	newPic()
}
func newPic() {
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
	paramMap, _ := JsonToMap(storeInfoStr, 1594137931634)
	// fmt.Printf("paramMap：：：：===》 %+v\n", paramMap)
	// func GenerateCert(tmplFilePath string, params maptest[string]interfacetest{}, w io.Writer) error {
	var templatePath = "E:\\20.06.16Project\\GoTest\\src\\MyTest\\writeTextToPic\\picandqrcode1\\demo\\delivery88.jpg"
	GenerateCert(templatePath, paramMap)
}
func main5() {
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
	paramMap, _ := JsonToMap(storeInfoStr, 1594137931634)
	fmt.Printf("paramMap：：：：===》 %+v\n", paramMap)

	// file := readTemplateFile("delivery.toml")

	// 从json模板中读取数据
	// template := GetPicTemplate("delivery.json")
	//
	// // dc := writeTextToPic(file, template, paramMap)
	// dc := writeTextToPic(template, paramMap)

	// bgImg, _ := gg.LoadImage(picPath)

	// 03: 生成二维码
	qrSize := 290
	content := "file.QrCode.Address"
	qrCodeImg, err = createAvatar(content, qrSize)
	if err != nil {
		fmt.Println("生成二维码失败:", err)
		return
	}
	// dc.DrawImage(qrCodeImg, 2551-458-qrSize, 3437-597-qrSize)
	// dc.SavePNG("123456789.png")

}

func JsonToMap(jsonStr string, onchianTime int64) (f map[string]interface{}, err error) {

	var storeInfo StoreInfo
	var infoByte = []byte(jsonStr)
	err = json.Unmarshal(infoByte, &storeInfo)
	if err != nil {
		panic(err)
	}
	var fieldMap = make(map[string]interface{})

	fieldMap["digest"] = storeInfo.Check.Digest
	fieldMap["hashAlgo"] = storeInfo.Check.HashAlgo
	fieldMap["sign"] = storeInfo.Check.Sign
	fieldMap["address"] = storeInfo.Data.Address
	fieldMap["signoffTime"] = storeInfo.Data.SignoffTime
	fieldMap["method"] = storeInfo.Data.Method
	fieldMap["sender"] = storeInfo.Data.Sender
	fieldMap["idcard"] = storeInfo.Data.Idcard
	fieldMap["people"] = storeInfo.Data.People
	fieldMap["sendTime"] = storeInfo.Data.SendTime
	fieldMap["bizSystemId"] = storeInfo.Header.BizSystemId
	fieldMap["caseId"] = storeInfo.Header.CaseId
	fieldMap["category"] = storeInfo.Header.Category
	fieldMap["courtId"] = storeInfo.Header.CourtId
	fieldMap["courtName"] = storeInfo.Header.CourtName
	fieldMap["subCategory"] = storeInfo.Header.SubCategory
	fieldMap["timestamp"] = storeInfo.Header.Timestamp
	fieldMap["onchianTime"] = onchianTime

	return fieldMap, nil
}

//  读取配置文件
func readTemplateFile(templateName string) (conf PicConfig) {
	var templateUrl = "D:/200707Go/gotest/src/MyTest/writeTextToPic/picandqrcode1/conf/" + templateName

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
func writeTextToPic( /*conf PicConfig,*/ template PicTemplate, param map[string]interface{}) *gg.Context {
	// 加载模板图片
	// templatePicPath := conf.ImageFilePath
	templatePicPath := template.ImagePath
	im, err := gg.LoadImage(templatePicPath)
	if err != nil {
		log.Fatal(err)
	}
	// 获取模板图片的尺寸
	width := im.Bounds().Size().X
	height := im.Bounds().Size().Y

	dc := gg.NewContextForImage(im)

	// dc.DrawImage(im, 0, 0)
	dc.DrawRoundedRectangle(0, 0, float64(width), float64(height), 0)
	// 将字体写入文件
	drawTextToPic(dc, template, param)

	// // bb.json、写主标题
	// titleFontFilePath := conf.TitleFontFilePath
	// mainTitleFontSize := conf.Text.MainTitleFontSize
	// mainTitleFontColor := conf.Text.MainTitleFontColor
	// mainTitleContent := conf.Text.MainTitleContent
	// titleStartPoint := conf.Text.MainTitleStartPoint
	// if err := dc.LoadFontFace(titleFontFilePath, float64(mainTitleFontSize)); err != nil {
	//	panic(err)
	// }
	// // 设置字体颜色
	// dc.SetColor(color.RGBA{R: uint8(mainTitleFontColor[0]), G: uint8(mainTitleFontColor[bb.json]), B: uint8(mainTitleFontColor[type2.toml]), A: 255})
	//
	// dc.DrawRoundedRectangle(0, 0, float64(width), float64(height), 0)
	//
	// titleCount := countNum(mainTitleContent)
	// dc.DrawStringAnchored(mainTitleContent, float64(titleStartPoint[0])+titleCount.word*float64(mainTitleFontSize)/4+titleCount.num*float64(mainTitleFontSize)/4, float64(titleStartPoint[bb.json]), 0.5, 0.5)
	//
	// // type2.toml、写入副标题
	// bodyFontFilePath := conf.BodyFontFilePath
	// subTitleFontSize := conf.Text.SubTitleFontSize
	// subTitleFontColor := conf.Text.SubTitleFontColor
	// subTitleContent := conf.Text.SubTitleContent
	// subTitleStartPoint := conf.Text.SubTitleStartPoint
	// if err := dc.LoadFontFace(bodyFontFilePath, float64(subTitleFontSize)); err != nil {
	//	panic(err)
	// }
	// // 设置字体颜色
	// dc.SetColor(color.RGBA{R: uint8(subTitleFontColor[0]), G: uint8(subTitleFontColor[bb.json]), B: uint8(subTitleFontColor[type2.toml]), A: 255})
	// subTitleCount := countNum(subTitleContent)
	// dc.DrawStringAnchored(subTitleContent, float64(subTitleStartPoint[0])+subTitleCount.word*float64(subTitleFontSize)/4+subTitleCount.num*float64(subTitleFontSize)/4, float64(subTitleStartPoint[bb.json]), 0.5, 0.5)
	//
	// // 发送人
	// sender := "相城法院"
	// // 当事人
	// people := "张三"
	// // 身份证号
	// idcard := "345678199303264569"
	// // 送达时间
	// time := "2020年06月11日 10:32:19"
	// // 上链时间
	// signoffTime := "2020年06月11日 10:32:26"
	// //  送达方式
	// category := "短信"
	//
	// bh := "TJ-XSSDFD123"
	//
	// var fontSize float64 = 90
	//
	// bhCount := countNum(bh)
	// dc.DrawStringAnchored(bh, 1140+bhCount.word*fontSize/4+bhCount.num*fontSize/4, 780, 0.5, 0.5)
	//
	// // 设置字体颜色
	// dc.SetColor(color.RGBA{A: 255})
	//
	// senderCount := countNum(sender)
	// dc.DrawStringAnchored(sender, 900+senderCount.hzc*fontSize/type2.toml+senderCount.num*fontSize/4, 679+360+20, 0.5, 0.5)
	//
	// peopleCount := countNum(people)
	// dc.DrawStringAnchored(people, 900+peopleCount.hzc*fontSize/type2.toml+peopleCount.num*fontSize/4, 1039+180+10+10, 0.5, 0.5)
	//
	// idcardCount := countNum(idcard)
	// dc.DrawStringAnchored(idcard, 900+idcardCount.word*fontSize/4+idcardCount.num*fontSize/4, 1419+10, 0.5, 0.5)
	//
	// categoryCount := countNum(category)
	// dc.DrawStringAnchored(category, 900+categoryCount.hzc*fontSize/type2.toml+categoryCount.num*fontSize/4, 1599+10, 0.5, 0.5)
	//
	// timeCount := countNum(time)
	// dc.DrawStringAnchored(time, 900+timeCount.hzc*fontSize/type2.toml+timeCount.num*fontSize/4+60, 1779+10, 0.5, 0.5)
	//
	// dc.DrawStringAnchored("同济区块链", 900+5*fontSize/type2.toml, 1959+10, 0.5, 0.5)
	//
	// signoffTimeCount := countNum(signoffTime)
	// dc.DrawStringAnchored(signoffTime, 900+signoffTimeCount.hzc*fontSize/type2.toml+signoffTimeCount.num*fontSize/4+60, 2139+10, 0.5, 0.5)

	dc.Clip()

	return dc
}

// 将字体写入文件
func drawTextToPic(dc *gg.Context, template PicTemplate, param map[string]interface{}) {

	// 获取默认字体路径
	defaultFontPath := template.DefaultFontPath
	// 获取默认字体大小
	fontDPI := template.FontDPI
	// 加载默认字体
	dc.LoadFontFace(defaultFontPath, fontDPI)

	// 获取blocks
	blocks := template.Blocks
	for _, blk := range blocks {
		blockType := blk.BlockType
		face := blk.FontFace
		size := blk.FontSize
		if len(face) >= 0 && size > 0 {
			dc.LoadFontFace(face, size)
		}
		dc.SetColor(color.RGBA{R: blk.Color[0], G: blk.Color[1], B: blk.Color[2], A: blk.Color[3]})

		switch blockType {
		case string(BlockType_Text):
			dc.DrawStringAnchored(blk.Content, blk.Point[0], blk.Point[1], 0, 0)
		case string(BlockType_DynamicText):
			var temp float64 = 0
			for _, arg := range blk.ContentList {
				fmt.Printf("arg:%+v\n", arg)
				value, has := param[arg.Field]
				if !has {
					continue
				}
				dc.DrawStringAnchored(arg.ShowValue, blk.Point[0], blk.Point[1]+temp, 0, 0)
				dc.DrawStringAnchored(fmt.Sprintf("%v", value), blk.Point[0]+1000, blk.Point[1]+1000+temp, 0, 0)
				temp += 300
			}
		}
	}

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
